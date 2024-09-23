package routes

import (
	"github.com/LukaMijovic/role-mgmt-access-ctrl/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	adminRoutes := server.Group("/admin")
	adminRoutes.POST("/login", loginAdmin)

	userRoutes := server.Group("/user")
	userRoutes.POST("/create", createUser)
	userRoutes.POST("/register", registerUser)
	userRoutes.POST("/login", loginUser)

	deviceRoutes := server.Group("/device")
	deviceRoutes.Use(middleware.Authenticate)
	deviceRoutes.POST("/register", registerDevice)
	deviceRoutes.POST("/unlock/:id", unlockRoom)

	accessRoutes := server.Group("/access")
	accessRoutes.Use(middleware.Authenticate)
	accessRoutes.POST("/temp/user/:id", receiveTempAccess)
}
