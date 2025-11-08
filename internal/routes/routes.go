package routes

import (
	"saint-seiya-awakening/internal/controllers"
	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("api/v1/register", controllers.RegisterUser)

	user := router.Group("/api/v1")
	{
		user.GET("/knights", controller.getKnights)
		user.GET("/knights/:id", controller.getKnightById)
	}

	admin := router.Group("/api/v1/admin", middleware.AuthMiddleware())
	{
		admin.POST("/knights", controller.createKnight)
	}
}