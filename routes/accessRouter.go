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

func receiveTempAccess(ctx *gin.Context) {
	accessId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, "Invalid url. Could not parse.")
	}

	var tempAccess model.TempAccessRight
	err = ctx.ShouldBindJSON(&tempAccess)
	//fmt.Printf("Access name: %v\n, endAccessDate: %v\n", tempAccess.GetAccessName(), tempAccess.GetEndAccessDate())

	if err != nil {
		//fmt.Printf("Error: %v\n", err.Error())
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Request body is invalid. Could not parse data")

		return
	}

	tempAccess.SetAccessId(accessId)

	//Admin signal for approval
	services.ConfirmTempRightByAdmin(&tempAccess)

	err = services.GiveTempAccessRightToUser(&tempAccess)

	_, err = services.LogRequest(tempAccess.UserID, "request for temp access")

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Could not save temp right.")

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessName": tempAccess.GetAccessName(),
		"message":    "Temp access granted",
		"validUntil": tempAccess.GetEndAccessDate(),
	})
}

func unlockRoom(ctx *gin.Context) {
	accessId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, "Invalid url. Could not parse.")

		return
	}

	//fmt.Printf("AccessRouter:\n accessId: %v\n", accessId)

	var tokenDTO dto.UserAccessDTO
	err = ctx.ShouldBindJSON(&tokenDTO)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Request body is invalid. Could not parse data.")

		return
	}

	id, ok := ctx.Get("userId")
	//fmt.Printf("UserId of logged in user from cookie: %v\n", id)

	if !ok {
		errorhandler.BadRequestError(ctx.JSON, http.StatusInternalServerError, "Something went wrong.")

		return
	}

	userId, ok := id.(int64)

	if !ok {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Invalid userId provided.")

		return
	}

	//fmt.Printf("UserId of logged in user parsed: %v\n", userId)

	if int64(userId) != tokenDTO.UserId {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Invalid userId.")

		return
	}

	deviceId, ok, err := services.CheckDeviceIMEIofUser(tokenDTO.IMEI, int64(userId))

	if err != nil {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, err.Error())

		return
	}

	if !ok {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, "There are no IMEIs found for provided user.")

		return
	}

	//fmt.Printf("AccessRouter:\n DeviceId: %v\n", deviceId)

	//log event deviceId
	accessIdParsed := fmt.Sprint(accessId)
	logId, err := services.LogEvent(int64(userId), int64(deviceId), "access "+accessIdParsed)

	if err != nil {
		errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Could not save access log.")

		return
	}

	//Access right check
	ok, err = services.CheckAccessRightOfUser(int64(userId), accessId)

	if err != nil {
		errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Could not get access rights from database.")

		return
	}

	if !ok {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, "User does not have needed rights to unlock the room.")
		//fmt.Printf("AccessRouter:\n UserId: %v, AccessId: %v, IMEI: %v\n", userId, accessId, tokenDTO.IMEI)

		return
	}

	unlockTime := time.Now()
	err = services.UpdateUnlockTime(logId, unlockTime, accessId)

	if err != nil {
		errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Could not update the unlock time")

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"eventId":    logId,
		"IMEI":       tokenDTO.IMEI,
		"message":    "OK to unlock the lock",
		"lockId":     accessId,
		"userId":     userId,
		"unlockTime": unlockTime,
	})
}

func lockRoom(ctx *gin.Context) {
	accessId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, "Invalid url. Could not parse.")

		return
	}

	var tokenDTO dto.UserAccessDTO
	err = ctx.ShouldBindJSON(&tokenDTO)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Invalid body. Could not parse data.")

		return
	}

	id, ok := ctx.Get("userId")
	// fmt.Printf("UserId of logged in user: %v\n", id)

	if !ok {
		errorhandler.BadRequestError(ctx.JSON, http.StatusInternalServerError, "Something went wrong.")

		return
	}

	userId, ok := id.(int64)
	//fmt.Printf("UserId of logged in user: %v\n", userId)

	if !ok {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Invalid userId provided.")

		return
	}

	if int64(userId) != tokenDTO.UserId {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Invalid userId.")

		return
	}

	deviceId, ok, err := services.CheckDeviceIMEIofUser(tokenDTO.IMEI, int64(userId))
	//fmt.Printf("DeviceId: %v\n", deviceId)

	if err != nil {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, err.Error())

		return
	}

	if !ok {
		errorhandler.BadRequestError(ctx.JSON, http.StatusBadRequest, "There are no IMEIs found for provided user.")

		return
	}

	logId, err := services.LogEvent(int64(userId), int64(deviceId), "lock "+fmt.Sprint(accessId))

	if err != nil {
		errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Could not save access log.")

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"eventId": logId,
		"IMEI":    tokenDTO.IMEI,
		"message": "Locked",
		"lockId":  accessId,
		"userId":  userId,
	})
}
