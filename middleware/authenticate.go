package middleware

import (
	"fmt"
	"net/http"
	"strings"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/util"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	fmt.Printf("Token received: %v\n", token)

	if token == "" {
		errorhandler.AuthenticationError(ctx, http.StatusUnauthorized, "Authorization header empty.")

		return
	}

	token, ok := strings.CutPrefix(token, "Bearer ")

	if !ok {
		errorhandler.AuthenticationError(ctx, http.StatusUnauthorized, "Invalid token received")

		return
	}

	userId, err := util.VerifyToken(token)

	if err != nil {
		errorhandler.AuthenticationError(ctx, http.StatusUnauthorized, err.Error())

		return
	}

	ctx.Set("userID", userId)
	ctx.Next()
}
