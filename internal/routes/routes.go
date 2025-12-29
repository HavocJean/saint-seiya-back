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

	v1.GET("/teams/public", app.TeamController.GetPublicTeams)

	auth := v1.Group("")
	auth.Use(app.AuthMiddleware)
	{
		auth.GET("/profile", app.AuthController.GetProfileUser)

		auth.POST("/team", app.TeamController.CreateTeam)
		auth.POST("/team/add/:id", app.TeamController.AddKnightToTeam)
		auth.DELETE("/team/:id", app.TeamController.DeleteTeam)
		auth.DELETE("/team/:teamId/knight/:knightId", app.TeamController.DeleteTeamKnight)
	}

	adminAuth := v1.Group("/admin")
	adminAuth.Use(app.AdminMiddleware)
	{
		adminAuth.POST("/knights", app.KnightController.CreateKnight)
		adminAuth.POST("/knights/:id/skills", app.KnightController.CreateKnightSkill)
		adminAuth.POST("/cosmos", app.CosmoController.CreateCosmo)
	}
}
