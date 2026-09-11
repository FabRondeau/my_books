package router

import (
	"goapi/security"

	"github.com/gin-gonic/gin"
)

func PrivateRouter(engine *gin.Engine) *gin.RouterGroup {
	// authorized := security.Auth(engine.Handler())

	protected := engine.Group("/private")
	protected.Use(security.AuthMiddleware())
	return protected
}
