package setup

import (
	"saint-seiya-back/internal/bootstrap"
	"saint-seiya-back/internal/infrastructure/http/middleware"
	"saint-seiya-back/internal/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTestServer(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)

	app := bootstrap.BuildApp(db)

	router := gin.New()
	router.Use(middleware.RecoveryMiddleware())
	routes.SetupRoutes(router, app)

	return router
}
