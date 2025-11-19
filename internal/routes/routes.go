package routes

import (
	"saint-seiya-awakening/internal/controllers"
	"saint-seiya-awakening/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", controllers.RegisterUser)
		v1.POST("/login", controllers.LoginUser)

		v1.GET("/cosmos", controllers.GetCosmos)
		v1.GET("/cosmos/:id", controllers.GetCosmosById)

		v1.GET("/knights", controllers.GetKnights)
		v1.GET("/knights/:id", controllers.GetKnightById)

		// v1.POST("/build", controllers.CreateBuild)
	}

	userAuth := v1.Group("/", middleware.AuthJwtMiddleware())
	{
		userAuth.GET("/profile", controllers.GetUserProfile)
	}

	adminAuth := v1.Group("/admin", middleware.AdminAuthMiddleware())
	{
		adminAuth.POST("/knights", controllers.CreateKnight)
		adminAuth.POST("/knights/:id/skills", controllers.CreateKnightSkill)
		adminAuth.POST("/cosmos", controllers.CreateCosmo)
	}
}
