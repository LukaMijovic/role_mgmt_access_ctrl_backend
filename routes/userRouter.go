package routes

import (
	"fmt"
	"net/http"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/services"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var credentials dto.UserCredentialsDTO
	err := ctx.ShouldBindJSON(&credentials)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx, http.StatusBadRequest)

		return
	}

	err = services.RegisterUserToDatabase(&credentials)

	if err != nil {
		fmt.Println(err.Error())
		errorhandler.DatabaseError(ctx, http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": credentials.User_ID,
		"email":   credentials.Email,
	})
}

func CreateUser(ctx *gin.Context) {
	var userDTO model.User
	err := ctx.ShouldBindJSON(&userDTO)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx, http.StatusBadRequest)

		return
	}

	//fmt.Printf(userDTO.Firstname + " " + userDTO.Lastname + " " + userDTO.Email + "\n")

	user, err := services.SaveUserToDatabase(&userDTO)

	if err != nil {
		errorhandler.DatabaseError(ctx, http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"userID":          user.GetID(),
		"registraionTime": user.GetUserRegistrationDate(),
	})
}
