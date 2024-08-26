package main

import (
	"github.com/gin-gonic/gin"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/routes"
)

func main() {
	database.ConnectToDatabase()
	defer database.DisconnectDatabase()

	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
