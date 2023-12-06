package api

import (
	"io"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/db"
)

// HttpMessageGet godoc
// @Summary get a message
// @Schemes get a message and it's content / other data
// @Description deletes a message
// @Tags Message API
// @Accept json
// @Produce json
// @Success 200
// @Router /api/message/id/:messageKey [get]
func HttpMessageGet(c *gin.Context) {
	if c.Param("messageKey") == "" {
		c.JSON(400, gin.H{
			"message": "Missing messageId",
		})
		return
	}

	// check for permissions within postgres
	message := db.GetMessage(c.Param("messageKey"))
	c.JSON(200, gin.H{
		"message": message,
	})

}

// HttpMessageDelete godoc
// @Summary deletes a message
// @Schemes
// @Description deletes a message
// @Tags Message API
// @Accept json
// @Produce json
// @Success 200
// @Router /api/message/id/:messageKey [delete]
func HttpMessageDelete(c *gin.Context) {
	if c.Param("messageKey") == "" {
		c.JSON(400, gin.H{
			"message": "Missing messageId",
		})
		return
	}

	// check for permissions within postgres

	db.DeleteMessage(c.Param("messageKey"))
	c.JSON(200, gin.H{
		"message": "Message deleted",
	})
}

// HttpMessageRoom godoc
// @Summary send a message to a given room
// @Schemes
// @Description deletes a message
// @Tags Message API, Room API
// @Accept json
// @Produce json
// @Success 200
// @Router /api/room/:roomKey/message [post]
func HttpMessageRoom(c *gin.Context) {
	if c.Param("roomKey") == "" {
		c.JSON(400, gin.H{
			"message": "Missing roomKey",
		})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	content, err := jsonparser.GetString(body, "content")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid content",
		})
		return
	}

	message := db.InsertNewMessage(c.Param("roomKey"), 0, content)

	c.JSON(200, gin.H{
		"message": message,
	})
}
