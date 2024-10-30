package routes

import (
	"github.com/LukaMijovic/role-mgmt-access-ctrl/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	adminRoutes := server.Group("/admin")
	adminRoutes.POST("/login", loginAdmin)
	adminRoutes.GET("/connect", connectToWS)
	adminRoutes.Use(middleware.Authenticate)
	adminRoutes.GET("/user/:id", getUser)

	userRoutes := server.Group("/user")
	userRoutes.POST("/create", createUser)
	userRoutes.POST("/register", registerUser)
	userRoutes.POST("/login", loginUser)
	userRoutes.GET("/connect", connectUserToWS)

	deviceRoutes := server.Group("/device")
	deviceRoutes.Use(middleware.Authenticate)
	deviceRoutes.POST("/register", registerDevice)

	accessRoutes := server.Group("/access")
	accessRoutes.Use(middleware.Authenticate)
	accessRoutes.POST("/temp/user/:id", receiveTempAccess)
	accessRoutes.POST("/unlock/:id", unlockRoom)
	accessRoutes.POST("/lock/:id", lockRoom)
}
