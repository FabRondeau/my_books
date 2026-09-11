package router

import (
	"goapi/controller"

	"github.com/gin-gonic/gin"
)

// AuthRouter ajoute les routes d'authentification à l'instance existante de gin.Engine
func AuthRouter(engine *gin.Engine, authController *controller.AuthController) {
	authGroup := engine.Group("/auth")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)
}
