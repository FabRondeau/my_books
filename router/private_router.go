package router

import (
	"goapi/controller"
	"goapi/security"

	"github.com/gin-gonic/gin"
)

func PrivateRouter(engine *gin.Engine, authorController *controller.AuthorController) *gin.RouterGroup {
	// authorized := security.Auth(engine.Handler())

	protected := engine.Group("/private")
	protected.Use(security.AuthMiddleware())
	return protected
}
