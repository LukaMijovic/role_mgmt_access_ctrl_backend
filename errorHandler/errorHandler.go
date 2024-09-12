package errorhandler

import (
	"github.com/gin-gonic/gin"
)

func BadBodyRequestError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message, //"Request body is invalid. Could not parse data",
	})
}

func DatabaseError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message, //"Error while saving object to database"
	})
}

func AuthenticationError(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
	})
}
