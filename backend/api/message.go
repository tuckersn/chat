package api

import (
	"github.com/gin-gonic/gin"
)

func MessageRoutes(r *gin.Engine) {

	r.GET("/message/*id", func(c *gin.Context) {

	})

	r.POST("/message/*id", func(c *gin.Context) {
	})

	// handle all paths of /page under one handler for DELETE
	r.DELETE("/message/*id", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
