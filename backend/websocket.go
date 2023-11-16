package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WebsocketClient(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		time.Sleep(time.Second)
	}
}

func WebsocketGinRoutes(r *gin.Engine) {

	// websocket client
	r.GET("/ws", func(c *gin.Context) {
		WebsocketClient(c)
	})
}
