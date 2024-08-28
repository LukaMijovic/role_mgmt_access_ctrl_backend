package routes

import (
	"net/http"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/services"
	"github.com/gin-gonic/gin"
)

func LoginAdmin(ctx *gin.Context) {
	var credentials dto.AdminCredentialsDTO
	err := ctx.ShouldBindJSON(&credentials)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx, http.StatusBadRequest)

		return
	}

	u, err := services.SignInAdmin(&credentials)

	if err != nil {
		errorhandler.AuthenticationError(ctx, http.StatusUnauthorized, err.Error())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": u.User_ID,
		"token":   "TOKEN",
	})
}
