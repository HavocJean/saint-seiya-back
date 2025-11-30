package routes

import (
	"saint-seiya-back/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, app *bootstrap.AppContext) {

	v1 := router.Group("/api/v1")
	v1.POST("/register", app.AuthController.RegisterUser)
	v1.POST("/login", app.AuthController.LoginUser)

	auth := v1.Group("")
	auth.Use(app.AuthMiddleware)
	{

	}

	// v1 := router.Group("/api/v1")
	// {
	// 	v1.GET("/cosmos", controllers.GetCosmos)
	// 	v1.GET("/cosmos/:id", controllers.GetCosmosById)

	// 	v1.GET("/knights", controllers.GetKnights)
	// 	v1.GET("/knights/:id", controllers.GetKnightById)
	// }

	// userAuth := v1.Group("/", middleware.AuthJwtMiddleware())
	// {
	// 	userAuth.GET("/profile", controllers.GetUserProfile)

	// 	userAuth.POST("/team", controllers.CreateTeam)
	// 	userAuth.POST("/team/add/:id", controllers.AddKnightToTeam)
	// 	userAuth.DELETE("/team/:id", controllers.DeleteTeam)
	// 	userAuth.DELETE("/team/:teamId/knight/:knightId", controllers.DeleteTeamKnight)
	// }

	// adminAuth := v1.Group("/admin", middleware.AdminAuthMiddleware())
	// {
	// 	adminAuth.POST("/knights", controllers.CreateKnight)
	// 	adminAuth.POST("/knights/:id/skills", controllers.CreateKnightSkill)
	// 	adminAuth.POST("/cosmos", controllers.CreateCosmo)
	// }
}
