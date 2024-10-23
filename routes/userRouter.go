package routes

import (
	"net/http"
	"strconv"

	errorhandler "github.com/LukaMijovic/role-mgmt-access-ctrl/errorHandler"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/services"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/util"
	"github.com/gin-gonic/gin"
)

func loginUser(ctx *gin.Context) {
	var credentials dto.UserCredentialsDTO
	err := ctx.ShouldBindJSON(&credentials)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Request body is invalid. Could not parse data")

		return
	}

	u, err := services.SignInUser(&credentials)

	if err != nil {
		errorhandler.AuthenticationError(ctx.AbortWithStatusJSON, http.StatusUnauthorized, err.Error())

		return
	}

	token, err := util.GenerateToken(u.Email, u.User_ID)

	if err != nil {
		errorhandler.AuthenticationError(ctx.AbortWithStatusJSON, http.StatusInternalServerError, "Token could not be generated")

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": u.User_ID,
		"token":   token,
	})
}

func registerUser(ctx *gin.Context) {
	// var credentials dto.UserCredentialsDTO
	// err := ctx.ShouldBindJSON(&credentials)

	// if err != nil {
	// 	errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Request body is invalid. Could not parse data")

	// 	return
	// }
	userId := ctx.Query("user_id")
	email := ctx.Query("email")
	password := ctx.Query("password")

	userIdParsed, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Request body is invalid. Could not parse data")

		return
	}

	credentials := dto.UserCredentialsDTO{
		User_ID:  userIdParsed,
		Email:    email,
		Password: password,
	}

	//Admin signal for approval
	go services.ConfirmCreationByAdmin(&credentials, ctx)

	// err = services.RegisterUserToDatabase(&credentials)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Error while saving object to database")

	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"user_id": credentials.User_ID,
	// 	"email":   credentials.Email,
	// })

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": credentials.User_ID,
		"email":   credentials.Email,
		"message": "Registration request has been sent.",
	})
}

func createUser(ctx *gin.Context) {
	var userDTO model.User
	err := ctx.ShouldBindJSON(&userDTO)

	if err != nil {
		errorhandler.BadBodyRequestError(ctx.JSON, http.StatusBadRequest, "Request body is invalid. Could not parse data")

		return
	}

	//fmt.Printf(userDTO.Firstname + " " + userDTO.Lastname + " " + userDTO.Email + "\n")

	user, err := services.SaveUserToDatabase(&userDTO)

	if err != nil {
		errorhandler.DatabaseError(ctx.JSON, http.StatusInternalServerError, "Error while saving object to database")

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"userID":          user.GetID(),
		"registraionTime": user.GetUserRegistrationDate(),
	})
}
