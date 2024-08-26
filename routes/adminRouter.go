package routes

import (
	"fmt"
	"net/http"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/gin-gonic/gin"
)

func LoginAdmin(ctx *gin.Context) {
	var credentials dto.AdminCredentialsDTO
	err := ctx.ShouldBindJSON(&credentials)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx, http.StatusBadRequest)

		return
	}

	//Logika za servis logovanja admina
	//...
	fmt.Printf("User %v, logged in with creds: %v\n", credentials.Email, credentials.Password)
}
