package errorhandler

import (
	"github.com/gin-gonic/gin"
)

func BadBodyRequestError(ctx *gin.Context, statusCode int) {
	ctx.JSON(statusCode, gin.H{
		"message": "Request body is invalid. Could not parse data",
	})
}
