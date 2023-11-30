package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tuckersn/chatbackend/api"
	_ "github.com/tuckersn/chatbackend/docs"
	docs "github.com/tuckersn/chatbackend/docs"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func httpServer() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	pagesDir := os.Getenv("COVE_PAGES_DIR")
	if pagesDir == "" {
		pagesDir = "~/.cove"
	}

	WebsocketGinRoutes(r)

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
			messageRouter.POST("/id/*messageId", api.HttpMessageSend)
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

	// api.SettingsRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on
}
