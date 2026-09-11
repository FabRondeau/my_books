package router

import (
	"goapi/controller"

	"github.com/gin-gonic/gin"
)

// PublicRouter ajoute les routes publiques à l'instance existante de gin.Engine
func PublicBookRouter(engine *gin.Engine, bookController *controller.BookController) {
	booksGroup := engine.Group("/books")
	booksGroup.GET("/all", bookController.DisplayAll)
	// booksGroup.POST("/add", bookController.Add)
}

func PrivateBookRouter(engine *gin.Engine, bookController *controller.BookController) {
	// bookized := security.Auth(engine.Handler())

	protected := PrivateRouter(engine)
	// protected := engine.Group("/private")
	// protected.Use(security.AuthMiddleware())
	privateBooksGroup := protected.Group("/book")
	privateBooksGroup.GET("/all", bookController.DisplayAll)
	privateBooksGroup.POST("/add", bookController.Add)

}
