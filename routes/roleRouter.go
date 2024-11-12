package routes

import (
	"net/http"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/services"
	"github.com/gin-gonic/gin"
)

func getAllRoles(ctx *gin.Context) {

	roles, err := services.GetAllRoles()

	if err != nil {
		errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Could not fetch roles.")

		return
	}

	ctx.JSON(http.StatusOK, roles)

}
