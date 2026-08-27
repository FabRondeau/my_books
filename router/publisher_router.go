package router

import (
	"goapi/controller"

	"github.com/gin-gonic/gin"
)

// PublicRouter ajoute les routes publiques à l'instance existante de gin.Engine
func PublisherRouter(engine *gin.Engine, publisherController *controller.PublisherController) {
	publishersGroup := engine.Group("/publishers")
	publishersGroup.GET("/all", publisherController.DisplayAll)
	publishersGroup.POST("/add", publisherController.Add)
}

// PublicRouter ajoute les routes publiques à l'instance existante de gin.Engine
