package api

import (
	"io"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/db"
)

// HttpMessageGet godoc
// @Summary Get a Message
// @Schemes
// @Description Get the Message with the given id (messageKey)
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
// @Summary Deletes a Message
// @Schemes
// @Description Deletes a Message
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
// @Summary Message a Room
// @Schemes
// @Description Send a Message to a given Room with the given content
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
