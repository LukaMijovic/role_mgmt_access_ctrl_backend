package routes

import (
	"net/http"
	"time"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/services"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func connectToWS(ctx *gin.Context) {
	wsHandler := &util.WebSocketHandler{
		websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	conn, err := wsHandler.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		//fmt.Printf("Error while upgrading from HTTP to WS: %v", err.Error())
		errorhandler.WebSocketConnectionError(ctx.JSON, http.StatusNotAcceptable, "Error while establishing Web Socket connection to the server.")
	}

	defer conn.Close()

	//Channel communication with Frontend Web application
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("hello, WebSocket!"))
		time.Sleep(time.Second * time.Duration(5))
	}
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
