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

type PublisherController struct {
	Db       *gorm.DB
	Validate *validator.Validate
}

func NewPublisherControllerImpl(Db *gorm.DB, validate *validator.Validate) *PublisherController {
	return &PublisherController{Db: Db, Validate: validate}
}

func (c PublisherController) DisplayAll(ctx *gin.Context) {

	var existingPublishers model.Publisher
	c.Db.Find(&existingPublishers)
	// fmt.Println(existingPublishers)
	ctx.JSON(http.StatusOK, existingPublishers)
}

func (c PublisherController) Add(ctx *gin.Context) {
	var reqBody request.PublisherNameRequest

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

	// var existingPublisher model.Publisher
	existingPublisher := c.FindByName(reqBody.Name)
	// fmt.Println(existingPublisher)
	if existingPublisher.Id > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Publisher " + existingPublisher.Name + " already exists"})
		return
	}

	// We can create the new publisher
	newPublisher := model.Publisher{
		Name: reqBody.Name,
	}

	if err := c.Db.Create(&newPublisher).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the publisher"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The publisher " + reqBody.Name + " registered successfully"})
}

func (c PublisherController) FindByName(name string) model.Publisher {

	var existingPublisher model.Publisher
	result := c.Db.Where("name = ?", name).First(&existingPublisher)

	fmt.Println(result.RowsAffected)
	if result.RowsAffected > 0 {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Publisher already exists"})

		return existingPublisher
	}
	fmt.Println(c.Db.Error)
	return model.Publisher{}
}
