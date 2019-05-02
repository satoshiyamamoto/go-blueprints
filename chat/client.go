package main

import "github.com/gorilla/websocket"

// clientはチャットを行なっている1人のユーザーを表します。
type client struct {
	socket *websocket.Conn // socketはこのクライアントのためのWebSocketです。
	send   chan []byte     // sendはメッセージが送られるチャネル。
	room   *room           // roomはこのクライアントが参加しているチャットルームです。
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
