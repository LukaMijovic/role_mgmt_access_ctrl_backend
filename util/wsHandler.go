package util

import "github.com/gorilla/websocket"

type WebSocketHandler struct {
	Upgrader websocket.Upgrader
}
