package errorhandler

import (
	"github.com/gin-gonic/gin"
)

func BadBodyRequestError(action func(int, interface{}), statusCode int, message string) {
	action(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message, //"Request body is invalid. Could not parse data",
	})
}

func BadRequestError(action func(int, interface{}), statusCode int, message string) {
	action(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
	})
}

func DatabaseError(action func(int, interface{}), statusCode int, message string) {
	action(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message, //"Error while saving object to database"
	})
}

func AuthenticationError(action func(int, interface{}), statusCode int, message string) {
	action(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
	})
}

func WebSocketConnectionError(action func(int, interface{}), statusCode int, message string) {
	action(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
	})
}
