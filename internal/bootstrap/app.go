package bootstrap

import (
	"saint-seiya-back/internal/application/auth"
	"saint-seiya-back/internal/application/cosmo"
	"saint-seiya-back/internal/application/knight"
	"saint-seiya-back/internal/application/team"
	"saint-seiya-back/internal/config"
	"saint-seiya-back/internal/infrastructure/database"
	"saint-seiya-back/internal/infrastructure/database/repositories"
	"saint-seiya-back/internal/infrastructure/http/controllers"
	"saint-seiya-back/internal/infrastructure/http/middleware"
	"saint-seiya-back/internal/infrastructure/services"
	"sync"

	"github.com/gin-gonic/gin"
)

var once sync.Once
var appCtxInstance *AppContext

type AppContext struct {
	AuthController   *controllers.AuthController
	KnightController *controllers.KnightController
	CosmoController  *controllers.CosmoController
	TeamController   *controllers.TeamController
	AuthMiddleware   gin.HandlerFunc
	AdminMiddleware  gin.HandlerFunc
}

func InitApp() *AppContext {
	once.Do(func() {
		db := database.GetDB()

		jwtService := services.NewJwtService(config.Cfg.JWTSecret)

		// repositories
		userRepository := repositories.NewUserRepository(db)
		knightRepository := repositories.NewKnightRepository(db)
		cosmoRepository := repositories.NewCosmoRepository(db)
		teamRepository := repositories.NewTeamRepository(db)

		// use cases
		loginUseCase := auth.NewLoginUseCase(userRepository, jwtService)
		registerUseCase := auth.NewRegisterUseCase(userRepository, jwtService)
		getUserByIdUseCase := auth.NewUserByIdUseCase(userRepository)

		createKnightUseCase := knight.NewCreateKnightUseCase(knightRepository)
		getKnightsUseCase := knight.NewGetKnightsUseCase(knightRepository)
		getKnightByIdUseCase := knight.NewGetKnightByIdUseCase(knightRepository)
		createKnightSkillUseCase := knight.NewCreateKnightSkillUseCase(knightRepository)

		getCosmosUseCase := cosmo.NewGetCosmosUseCase(cosmoRepository)
		getCosmoByIdUseCase := cosmo.NewGetCosmoByIdUseCase(cosmoRepository)
		createCosmoUseCase := cosmo.NewCreateCosmoUseCase(cosmoRepository)

		createTeamUseCase := team.NewCreateTeamUseCase(teamRepository)
		addKnightToTeamUseCase := team.NewAddKnightToTeamUseCase(teamRepository)
		deleteKnightToTeamUseCase := team.NewDeleteKnightToTeamUseCase(teamRepository)
		deleteTeamUseCase := team.NewDeleteTeamUseCase(teamRepository)
		getPublicTeamsUseCase := team.NewGetPublicTeamsUseCase(teamRepository)

		// controllers
		authController := controllers.NewAuthController(loginUseCase, registerUseCase, getUserByIdUseCase)
		knightController := controllers.NewKnightController(
			createKnightUseCase,
			getKnightsUseCase,
			getKnightByIdUseCase,
			createKnightSkillUseCase,
		)

		cosmoController := controllers.NewCosmoController(
			getCosmosUseCase,
			getCosmoByIdUseCase,
			createCosmoUseCase,
		)

		teamController := controllers.NewTeamController(
			createTeamUseCase,
			addKnightToTeamUseCase,
			deleteTeamUseCase,
			deleteKnightToTeamUseCase,
			getPublicTeamsUseCase,
		)

		// middleware and instance
		authMiddleware := middleware.AuthJwtMiddleware(jwtService)
		adminMiddleware := middleware.AdminAuthMiddleware()

		appCtxInstance = &AppContext{
			AuthController:   authController,
			KnightController: knightController,
			CosmoController:  cosmoController,
			TeamController:   teamController,
			AuthMiddleware:   authMiddleware,
			AdminMiddleware:  adminMiddleware,
		}
	})

	return appCtxInstance
}
