package router

import (
	"goapi/controller"

	"github.com/gin-gonic/gin"
)

// PublicRouter ajoute les routes publiques à l'instance existante de gin.Engine
func PublicAuthorRouter(engine *gin.Engine, authorController *controller.AuthorController) {
	authorsGroup := engine.Group("/authors")
	authorsGroup.GET("/all", authorController.DisplayAll)
	// authorsGroup.POST("/add", authorController.Add)
}

func PrivateAuthorRouter(engine *gin.Engine, authorController *controller.AuthorController) {
	// authorized := security.Auth(engine.Handler())

	protected := PrivateRouter(engine, authorController)
	// protected := engine.Group("/private")
	// protected.Use(security.AuthMiddleware())
	privateAuthorsGroup := protected.Group("/author")
	privateAuthorsGroup.GET("/all", authorController.DisplayAll)
	privateAuthorsGroup.POST("/add", authorController.Add)

}
