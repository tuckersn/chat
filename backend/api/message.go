package api

import (
	"github.com/gin-gonic/gin"
)

// HttpMessageGet godoc
// @Summary get a message
// @Schemes get a message and it's content / other data
// @Description deletes a message
// @Tags Message API
// @Accept json
// @Produce json
// @Success 200
// @Router /api/message/id/:messageId [get]
func HttpMessageGet(c *gin.Context) {
}

// HttpMessageDelete godoc
// @Summary deletes a message
// @Schemes
// @Description deletes a message
// @Tags Message API
// @Accept json
// @Produce json
// @Success 200
// @Router /api/message/id/:messageId [delete]
func HttpMessageDelete(c *gin.Context) {
}

// HttpMessageSend godoc
// @Summary send a message to a given room
// @Schemes
// @Description deletes a message
// @Tags Message API, Room API
// @Accept json
// @Produce json
// @Success 200
// @Router /api/room/:roomId/message [post]
func HttpMessageSend(c *gin.Context) {
}
