package routes

import (
	"fmt"
	"net/http"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/services"
	"github.com/gin-gonic/gin"
)

func registerDevice(ctx *gin.Context) {
	var deviceDTO model.Device
	err := ctx.ShouldBindJSON(&deviceDTO)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Request body is invalid. Could not parse data")

		return
	}

	device, err := services.SaveDeviceToDatabase(&deviceDTO)

	if err != nil {
		fmt.Println(err.Error())
		errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Error while saving object to database")

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"deviceID":         device.GetID(),
		"registrationTime": device.GetDeviceRegistraionDate(),
	})
}
