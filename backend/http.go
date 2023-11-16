package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tuckersn/chatbackend/api"
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
			pageRouter.GET("/*path", api.HttpGetPage)
			pageRouter.POST("/*path", api.HttpPostPage)
			pageRouter.DELETE("/*path", api.HttpDeletePage)
		}
	}

	api.UserRoutes(r)
	api.MessageRoutes(r)
	api.SettingsRoutes(r)
	api.ServerRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on
}
