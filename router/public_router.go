package router

import (
	"goapi/controller"

	"github.com/gin-gonic/gin"
)

// PublicRouter ajoute les routes publiques à l'instance existante de gin.Engine
func PublicRouter(engine *gin.Engine, publicController *controller.PublicController) {
	booksGroup := engine.Group("/books")
	booksGroup.GET("/all", publicController.DisplayAll)
}
