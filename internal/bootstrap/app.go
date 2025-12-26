package bootstrap

import (
	"saint-seiya-back/internal/application/auth"
	"saint-seiya-back/internal/application/cosmo"
	"saint-seiya-back/internal/application/knight"
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
	AuthMiddleware   gin.HandlerFunc
}

func InitApp() *AppContext {
	once.Do(func() {
		db := database.GetDB()

		jwtService := services.NewJwtService(config.Cfg.JWTSecret)

		userRepository := repositories.NewUserRepository(db)
		knightRepository := repositories.NewKnightRepository(db)
		cosmoRepository := repositories.NewCosmoRepository(db)

		loginUseCase := auth.NewLoginUseCase(userRepository, jwtService)
		registerUseCase := auth.NewRegisterUseCase(userRepository, jwtService)

		createKnightUseCase := knight.NewCreateKnightUseCase(knightRepository)
		getKnightsUseCase := knight.NewGetKnightsUseCase(knightRepository)
		getKnightByIdUseCase := knight.NewGetKnightByIdUseCase(knightRepository)
		getCosmosUseCase := cosmo.NewGetCosmosUseCase(cosmoRepository)
		getCosmoByIdUseCase := cosmo.NewGetCosmoByIdUseCase(cosmoRepository)

		authController := controllers.NewAuthController(loginUseCase, registerUseCase)
		knightController := controllers.NewKnightController(
			createKnightUseCase,
			getKnightsUseCase,
			getKnightByIdUseCase,
		)
		cosmoController := controllers.NewCosmoController(
			getCosmosUseCase,
			getCosmoByIdUseCase,
		)

		authMiddleware := middleware.AuthJwtMiddleware(jwtService)

		appCtxInstance = &AppContext{
			AuthController:   authController,
			KnightController: knightController,
			CosmoController:  cosmoController,
			AuthMiddleware:   authMiddleware,
		}
	})

	return appCtxInstance
}
