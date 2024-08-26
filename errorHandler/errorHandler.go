package errorhandler

import (
	"github.com/gin-gonic/gin"
)

func BadBodyRequestError(ctx *gin.Context, statusCode int) {
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    "Request body is invalid. Could not parse data",
	})
}

func DatabaseError(ctx *gin.Context, statusCode int) {
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    "Error while saving object to database",
	})
}
