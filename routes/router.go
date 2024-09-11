package routes

import (
	"github.com/LukaMijovic/role-mgmt-access-ctrl/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	adminRoutes := server.Group("/admin")
	adminRoutes.POST("/login", LoginAdmin)

	userRoutes := server.Group("/user")
	userRoutes.POST("/create", CreateUser)
	userRoutes.POST("/register", RegisterUser)

	deviceRoutes := server.Group("/device")
	deviceRoutes.POST("/register", RegisterDevice).Use(middleware.Authenticate)
}
