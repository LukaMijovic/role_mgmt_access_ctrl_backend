package routes

import (
	"net/http"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/gin-gonic/gin"
)

func loginAdmin(ctx *gin.Context) {
	var credentials dto.AdminCredentialsDTO
	err := ctx.ShouldBindJSON(&credentials)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Request body is invalid. Could not parse data",
		})

		return
	}

	//Logika za servis logovanja admina
	//...
}
