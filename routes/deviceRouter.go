package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
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

func unlockRoom(ctx *gin.Context) {
	accessId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	fmt.Printf("Room id: %v\n", accessId)

	if err != nil {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, "Invalid url. Could not parse.")

		return
	}

	var tokenDTO dto.UserAccessDTO
	err = ctx.ShouldBindJSON(&tokenDTO)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Request body is invalid. Could not parse data.")

		return
	}

	id, ok := ctx.Get("userId")
	// fmt.Printf("UserId of logged in user: %v\n", id)

	if !ok {
		errorhandler.BadRequestError(ctx.JSON, http.StatusInternalServerError, "Something went wrong.")

		return
	}

	userId, ok := id.(int64)
	fmt.Printf("UserId of logged in user: %v\n", userId)

	if !ok {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Invalid userId provided.")

		return
	}

	if int64(userId) != tokenDTO.UserId {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Invalid userId.")

		return
	}

	ok, err = services.CheckDeviceIMEIofUser(tokenDTO.IMEI, int64(userId))

	if err != nil {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, err.Error())

		return
	}

	if !ok {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, "There are no IMEIs found for provided user.")

		return
	}

	//Access right check

	ctx.JSON(http.StatusOK, gin.H{
		"IMEI":       tokenDTO.IMEI,
		"message":    "OK to unlock the lock",
		"lockId":     accessId,
		"userId":     userId,
		"unlockTime": time.Now(),
	})
}
