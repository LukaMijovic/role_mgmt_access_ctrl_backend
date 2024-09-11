package middleware

import (
	"net/http"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/util"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		errorhandler.AuthenticationError(ctx, http.StatusUnauthorized, "Authorization header empty.")

		return
	}

	userId, err := util.VerifyToken(token)

	if err != nil {
		errorhandler.AuthenticationError(ctx, http.StatusUnauthorized, "Invalid header token.")

		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
