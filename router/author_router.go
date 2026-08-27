package router

import (
	"goapi/controller"

	"github.com/gin-gonic/gin"
)

// PublicRouter ajoute les routes publiques à l'instance existante de gin.Engine
func AuthorRouter(engine *gin.Engine, authorController *controller.AuthorController) {
	authorsGroup := engine.Group("/authors")
	authorsGroup.GET("/all", authorController.DisplayAll)
	authorsGroup.POST("/add", authorController.Add)
}

// PublicRouter ajoute les routes publiques à l'instance existante de gin.Engine
