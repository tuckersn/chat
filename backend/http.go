package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tuckersn/chatbackend/api"
	docs "github.com/tuckersn/chatbackend/docs"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func httpServer(_db *sqlx.DB) {
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
			pageRouter.GET("/*path", api.HttpPageGet)
			pageRouter.POST("/*path", api.HttpPagePost)
			pageRouter.DELETE("/*path", api.HttpPageDelete)
		}
		messageRouter := apiRouter.Group("/message")
		{
			messageRouter.POST("/", api.HttpMessageCreate)
			messageRouter.GET("/*messageId", api.HttpMessageGet)
			messageRouter.POST("/*messageId", api.HttpMessageUpdate)
			messageRouter.DELETE("/*messageId", api.HttpMessageDelete)
		}
		userRouter := apiRouter.Group("/user")
		{
			userRouter.GET("/", api.HttpUserGet)
			userRouter.GET("/*userId", api.HttpUserGet)
			userRouter.POST("/", api.HttpUserCreate)
			userRouter.POST("/*userId", api.HttpUserUpdate)
			userRouter.GET("/*userId", api.HttpUserGet)
			userRouter.DELETE("/*userId", api.HttpUserDelete)
		}
	}

	api.SettingsRoutes(r)
	api.ServerRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on
}
