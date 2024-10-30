package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	Upgrader *websocket.Upgrader
}

var WebAppConnection *websocket.Conn
var MobileAppConnection *websocket.Conn

func (wsh *WebSocketHandler) Connect(ctx *gin.Context) (*websocket.Conn, error) {

	wsh.Upgrader.CheckOrigin = func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "http://localhost:4200"
		return true
	}

	conn, err := wsh.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		fmt.Printf("Error while upgrading from HTTP to WS: %v", err.Error())
		//errorhandler.WebSocketConnectionError(ctx.JSON, http.StatusNotAcceptable, "Error while establishing Web Socket connection to the server.")

		return nil, err
	}

	return conn, nil
}

func (wsh *WebSocketHandler) Disconnect(conn *websocket.Conn) {
	conn.Close()
}
