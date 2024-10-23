package util

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	Upgrader *websocket.Upgrader
}

var WebAppConnection *websocket.Conn

func (wsh *WebSocketHandler) Connect(ctx *gin.Context) (*websocket.Conn, error) {

	conn, err := wsh.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		//fmt.Printf("Error while upgrading from HTTP to WS: %v", err.Error())
		//errorhandler.WebSocketConnectionError(ctx.JSON, http.StatusNotAcceptable, "Error while establishing Web Socket connection to the server.")

		return nil, err
	}

	return conn, nil
}

func (wsh *WebSocketHandler) Disconnect(conn *websocket.Conn) {
	conn.Close()
}
