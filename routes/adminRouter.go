package routes

import (
	"fmt"
	"net/http"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/services"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func connectToWS(ctx *gin.Context) {
	wsHandler := &util.WebSocketHandler{
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	conn, err := wsHandler.Connect(ctx)

	if err != nil {
		fmt.Println(err.Error())
		errorhandler.WebSocketConnectionError(ctx.JSON, http.StatusNotAcceptable, "Error while establishing Web Socket connection to the server.")

		return
	}

	util.WebAppConnection = conn

	//Channel communication with Frontend Web application

	// for i := 0; i <= 10; i++ {
	// 	msg := "hello, WebSocket! " + fmt.Sprint(uint8(i))
	// 	conn.WriteMessage(websocket.TextMessage, []byte(msg))
	// 	time.Sleep(time.Second)
	// }

	//defer wsHandler.Disconnect(conn)
}

func loginAdmin(ctx *gin.Context) {
	var credentials dto.AdminCredentialsDTO
	err := ctx.ShouldBindJSON(&credentials)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Request body is invalid. Could not parse data")

		return
	}

	u, err := services.SignInAdmin(&credentials)

	if err != nil {
		errorhandler.AuthenticationError(ctx.AbortWithStatusJSON, http.StatusUnauthorized, err.Error())

		return
	}

	token, err := util.GenerateToken(u.Email, u.User_ID)

	if err != nil {
		errorhandler.AuthenticationError(ctx.AbortWithStatusJSON, http.StatusInternalServerError, "Token could not be generated")

		return
	}

	_, err = services.LogRequest(u.User_ID, "admin login")

	if err != nil {
		errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Could not save access log.")

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": u.User_ID,
		"token":   token,
	})
}
