package bootstrap

import (
	"saint-seiya-back/internal/application/auth"
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
	AuthController *controllers.AuthController
	AuthMiddleware gin.HandlerFunc
}

func InitApp() *AppContext {
	once.Do(func() {
		db := database.GetDB()

		jwtService := services.NewJwtService(config.Cfg.JWTSecret)

		userRepository := repositories.NewUserRepository(db)

		loginUseCase := auth.NewLoginUseCase(userRepository, jwtService)
		registerUseCase := auth.NewRegisterUseCase(userRepository, jwtService)

		authController := controllers.NewAuthController(loginUseCase, registerUseCase)
		authMiddleware := middleware.AuthJwtMiddleware(jwtService)

		appCtxInstance = &AppContext{
			AuthController: authController,
			AuthMiddleware: authMiddleware,
		}
	})

	return appCtxInstance
}
