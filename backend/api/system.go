package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

// HttpPing godoc
// @Summary ping the server, used for overview page
// @Schemes
// @Description says hello to the server, used for overview page
// @Tags System API
// @Accept json
// @Produce json
// @Success 200
// @Router /api/system/ping [get]
func HttpPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  "pong",
		"received": time.Now(),
	})
}
