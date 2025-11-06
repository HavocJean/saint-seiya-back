package routes

import (
	"saint-seiya-awakening/internal/controllers"
	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
}