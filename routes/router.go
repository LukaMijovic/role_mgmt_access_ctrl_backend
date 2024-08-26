package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	adminRoutes := server.Group("/admin")
	adminRoutes.POST("/login", LoginAdmin)
}
