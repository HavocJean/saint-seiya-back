package routes

import (
	"saint-seiya-awakening/internal/controllers"
	"saint-seiya-awakening/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("api/v1/register", controllers.RegisterUser)

	user := router.Group("/api/v1")
	{
		user.GET("/knights", controllers.GetKnights)
		user.GET("/knights/:id", controllers.GetKnightById)
	}

	admin := router.Group("/api/v1/admin", middleware.AdminAuthMiddleware())
	{
		admin.POST("/knights", controllers.CreateKnight)
		admin.POST("/cosmos", controllers.CreateCosmo)
	}
}
