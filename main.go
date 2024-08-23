package main

import (
	"github.com/gin-gonic/gin"

	"lukamijovic.com/role-mgmt-access-ctrl/database"
)

func main() {
	database.ConnectToDatabase()
	defer database.DisconnectDatabase()

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
