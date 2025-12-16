package main

import (
	"net/http"

	"github.com/Miklakapi/go-webosckets/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	hub := websocket.NewHub()
	go hub.Run()

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.GET("/ws", func(c *gin.Context) {
		websocket.ServeWS(c, hub)
	})

	engine.Run(":8000")
}
