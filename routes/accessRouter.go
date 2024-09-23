package routes

import (
	"fmt"
	"net/http"
	"strconv"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
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
