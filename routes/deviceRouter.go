package routes

import (
	"fmt"
	"net/http"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/services"
	"github.com/gin-gonic/gin"
)

func RegisterDevice(ctx *gin.Context) {
	var deviceDTO model.Device
	err := ctx.ShouldBindJSON(&deviceDTO)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx, http.StatusBadRequest)
	}

	device, err := services.SaveDeviceToDatabase(&deviceDTO)

	if err != nil {
		fmt.Println(err.Error())
		errorhandler.DatabaseError(ctx, http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"deviceID":         device.GetID(),
		"registrationTime": device.GetDeviceRegistraionDate(),
		"token":            "TOKEN",
	})
}
