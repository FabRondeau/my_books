package controller

import (
	"goapi/model"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PublicController struct {
	Db       *gorm.DB
	Validate *validator.Validate
}

func NewPublicControllerImpl(Db *gorm.DB, validate *validator.Validate) *PublicController {
	return &PublicController{Db: Db, Validate: validate}
}
func (c PublicController) DisplayAll(ctx *gin.Context) {

	var existingBooks model.Book
	result := c.Db.Find(&existingBooks)

	ctx.JSON(http.StatusOK, result)
}
