package routes

import (
	"saint-seiya-back/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, app *bootstrap.AppContext) {

	v1 := router.Group("/api/v1")

	v1.POST("/register", app.AuthController.RegisterUser)
	v1.POST("/login", app.AuthController.LoginUser)

	v1.GET("/knights", app.KnightController.GetKnights)
	v1.GET("/knights/:id", app.KnightController.GetKnightByID)

	v1.GET("/cosmos", app.CosmoController.GetCosmos)
	v1.GET("/cosmos/:id", app.CosmoController.GetCosmoById)

	auth := v1.Group("")
	auth.Use(app.AuthMiddleware)
	{
		// 	userAuth.GET("/profile", controllers.GetUserProfile)

		// 	userAuth.POST("/team", controllers.CreateTeam)
		// 	userAuth.POST("/team/add/:id", controllers.AddKnightToTeam)
		// 	userAuth.DELETE("/team/:id", controllers.DeleteTeam)
		// 	userAuth.DELETE("/team/:teamId/knight/:knightId", controllers.DeleteTeamKnight)
	}

	// adminAuth := v1.Group("/admin", middleware.AdminAuthMiddleware())
	// // {
	v1.POST("/knights", app.KnightController.CreateKnight)
	// adminAuth.POST("/knights/:id/skills", controllers.CreateKnightSkill)
	// 	adminAuth.POST("/cosmos", controllers.CreateCosmo)
	// }
}
