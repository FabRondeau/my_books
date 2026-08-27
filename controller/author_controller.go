package controller

import (
	"fmt"
	"goapi/data/request"
	"goapi/model"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthorController struct {
	Db       *gorm.DB
	Validate *validator.Validate
}

func NewAuthorControllerImpl(Db *gorm.DB, validate *validator.Validate) *AuthorController {
	return &AuthorController{Db: Db, Validate: validate}
}

func (c AuthorController) DisplayAll(ctx *gin.Context) {

	var existingAuthors model.Author
	c.Db.Find(&existingAuthors)
	// fmt.Println(existingAuthors)
	ctx.JSON(http.StatusOK, existingAuthors)
}

func (c AuthorController) Add(ctx *gin.Context) {
	var reqBody request.AuthorFullNameRequest

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	// var existingAuthor model.Author
	existingAuthor := c.FindByFullName(reqBody.FullName)
	// fmt.Println(existingAuthor)
	if existingAuthor.Id > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Author " + existingAuthor.FullName + " already exists"})
		return
	}

	// We can create the new author
	newAuthor := model.Author{
		FullName: reqBody.FullName,
	}

	if err := c.Db.Create(&newAuthor).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the author"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The author " + reqBody.FullName + " registered successfully"})
}

func (c AuthorController) FindByFullName(full_name string) model.Author {

	var existingAuthor model.Author
	result := c.Db.Where("full_name = ?", full_name).First(&existingAuthor)

	fmt.Println(result.RowsAffected)
	if result.RowsAffected > 0 {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Author already exists"})

		return existingAuthor
	}

	return model.Author{}
}
