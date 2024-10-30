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
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "X-Custom-Header", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowWebSockets = true
	//corsHandler := cors.Default()
	router.Use(cors.New(config))
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
