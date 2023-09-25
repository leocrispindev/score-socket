package service

import "github.com/gorilla/websocket"

var sessions map[int][]*websocket.Conn

func Add(id int, c *websocket.Conn) {
	conn := sessions[id]
	conn = append(conn, c)
}

func GetSubscribersByMatchID(id int) []*websocket.Conn {
	return sessions[id]
}
