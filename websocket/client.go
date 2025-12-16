package websocket

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pingPeriod     = (pongWait * 9) / 10
	pongWait       = 60 * time.Second
	writeWait      = 10 * time.Second
	maxMessageSize = 512
)

type Client struct {
	Conn *websocket.Conn
	send chan Message
	hub  *Hub
}

func NewClient(conn *websocket.Conn, hub *Hub) *Client {
	return &Client{Conn: conn, send: make(chan Message, 256), hub: hub}
}

func (c *Client) Close() {
	close(c.send)
}

func (c *Client) Read() {
	defer func() {
		c.hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				fmt.Println("Client close connection")
				return
			}

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println("Unexpected close:", err)
				return
			}

			fmt.Println("WS read error:", err)
			return
		}
		c.hub.broadcast <- Message{Sender: c, Data: msg}
	}
}

func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if ok {
				err := c.Conn.WriteMessage(websocket.TextMessage, message.Data)
				if err != nil {
					fmt.Println("Error: ", err)
					break
				}
			} else {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}

	}
}
