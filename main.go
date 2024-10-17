package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/routes"
)

func main() {
	database.ConnectToDatabase()
	defer database.DisconnectDatabase()

	router := gin.Default()
	//config := cors.DefaultConfig()
	router.Use(cors.Default())
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
