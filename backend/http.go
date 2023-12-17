package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tuckersn/chatbackend/api"
	_ "github.com/tuckersn/chatbackend/docs"
	docs "github.com/tuckersn/chatbackend/docs"
	"github.com/tuckersn/chatbackend/util"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func generateSelfSignedCert() {
	logger := log.New(os.Stdout, "[SELF CERTIFICATE]", log.LstdFlags|log.Lshortfile)

	// Generate private key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logger.Fatalf("Failed to generate private key: %v", err)
	}

	// Serialize private key to PEM format
	keyOut, err := os.Create(util.Config.Http.KeyFile)
	if err != nil {
		logger.Fatalf("Failed to open cert.key for writing: %v", err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()

	// Create a self-signed certificate
	notBefore := time.Now()
	notAfter := notBefore.Add(3650 * 24 * time.Hour)

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		logger.Fatalf("Failed to generate serial number: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Your Organization"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	template.IPAddresses = append(template.IPAddresses, net.ParseIP("127.0.0.1"))

	// Create certificate
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		logger.Fatalf("Failed to create certificate: %v", err)
	}

	// Serialize certificate to PEM format
	certOut, err := os.Create(util.Config.Http.CertFile)
	if err != nil {
		logger.Fatalf("Failed to open cert.crt for writing: %v", err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	logger.Printf("cert.key and cert.crt have been successfully generated")
}

func httpServer() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRouter := r.Group("/api")
	{
		pageRouter := apiRouter.Group("/page")
		{
			pageRouter.GET("/*path", api.HttpNoteGet)
			pageRouter.POST("/*path", api.HttpNotePost)
			pageRouter.DELETE("/*path", api.HttpNoteDelete)
		}
		messageRouter := apiRouter.Group("/message")
		{
			messageRouter.GET("/id/*messageId", api.HttpMessageGet)
			messageRouter.POST("/id/*messageId", api.HttpMessageRoom)
			messageRouter.DELETE("/id/*messageId", api.HttpMessageDelete)
		}
		userRouter := apiRouter.Group("/user")
		{
			userRouter.POST("/", api.HttpUserCreate)
			userRouter.GET("/id/*username", api.HttpUserGet)
			userRouter.POST("/id/*username", api.HttpUserUpdate)
			userRouter.DELETE("/id/*username", api.HttpUserDelete)
		}
		serverRouter := apiRouter.Group("/server")
		{
			serverRouter.GET("/ping", api.HttpPing)
		}
		webhookRouter := apiRouter.Group("/webhook")
		{
			webhookRouter.GET("/", api.HttpWebhookList)
			webhookRouter.POST("/", api.HttpWebhookCreate)
			webhookRouter.GET("/id/*webhookId", api.HttpWebhookGet)
			webhookRouter.POST("/id/*webhookId", api.HttpWebhookUpdate)
			webhookRouter.DELETE("/id/*webhookId", api.HttpWebhookDelete)
		}
	}

	loginRouter := r.Group("/login")
	{
		loginRouter.GET("/recent", api.HttpLoginRecent)
		if util.Config.GitLab.AuthEnabled {
			loginRouter.GET("/gitlab", api.HttpLoginGitlabRedirect)
			loginRouter.GET("/gitlab/receive", api.HttpLoginGitlabReceive)
		}
		if util.Config.Google.AuthEnabled {
			loginRouter.GET("/google", api.HttpLoginGoogle)
			loginRouter.GET("/google/receive", api.HttpLoginGoogleReceiveToken)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if _, err := os.Stat(util.Config.Http.CertFile); os.IsNotExist(err) {
		generateSelfSignedCert()
	}
	r.RunTLS(":"+fmt.Sprint(util.Config.Http.Port), util.Config.Http.CertFile, util.Config.Http.KeyFile)
}
