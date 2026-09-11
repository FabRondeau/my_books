package controller

import (
	"encoding/json"
	"fmt"
	"goapi/data/request"
	"goapi/model"
	"log"
	"net/http"
	"net/url"
	"os"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookController struct {
	Db       *gorm.DB
	Validate *validator.Validate
}

func NewBookControllerImpl(Db *gorm.DB, validate *validator.Validate) *BookController {
	return &BookController{Db: Db, Validate: validate}
}

func (c BookController) DisplayAll(ctx *gin.Context) {

	var existingBooks []model.Book
	c.Db.Find(&existingBooks)
	ctx.JSON(http.StatusOK, existingBooks)
}

func (c BookController) Add(ctx *gin.Context) {
	var reqBody request.BookISBN13Request

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

	// var existingBook model.Book
	existingBook := c.FindByISBN13(reqBody.ISBN13)
	if existingBook.Id > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book " + existingBook.ISBN13 + " already exists"})
		return
	}

	// We can create the new author
	newBook := model.Book{
		ISBN13: reqBody.ISBN13,
	}

	googleBook := c.FindGoogleBook(reqBody.ISBN13)
	fmt.Println(googleBook)
	if googleBook.TotalItems == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book " + reqBody.ISBN13 + " not found in Google books."})
		return
	}

	// Résultat: Google book trouvé:on crée les auteurs et l'éditeur s'ils n'existent pas

	fmt.Println(googleBook.TotalItems)

	// On prend le premier item
	googleBookItem := googleBook.Items[0]
	fmt.Println(googleBookItem)

	// Créer l'éditeur s'il n'existe pas

	// Créé

	if err := c.Db.Create(&newBook).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the author"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "The author " + reqBody.ISBN13 + " registered successfully"})
}

func (c BookController) FindByISBN13(isbn13 string) model.Book {

	var existingBook model.Book
	result := c.Db.Where("ISBN13 = ?", isbn13).First(&existingBook)

	if result.RowsAffected > 0 {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book already exists"})
		return existingBook
	}

	return model.Book{}
}

func (c BookController) FindGoogleBook(isbn13 string) request.GoogleBookRequest {
	t := http.DefaultTransport.(*http.Transport).Clone()

	// Use either HTTP/1 and HTTP/2.
	t.Protocols = new(http.Protocols)
	t.Protocols.SetHTTP1(true)
	t.Protocols.SetHTTP2(true)
	client := &http.Client{Transport: t}
	googleBooksBaseUrl := os.Getenv("GOOGLEBOOKS_VOLUMES_URL")
	googleBooksKey := os.Getenv("GOOGLEBOOKS_KEY")
	endpoint, err := url.Parse(googleBooksBaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	queryParams := url.Values{}
	queryParams.Set("key", googleBooksKey)
	queryParams.Set("q", "isbn:"+isbn13)
	endpoint.RawQuery = queryParams.Encode()

	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	decoder := json.NewDecoder(response.Body)
	var googleBook request.GoogleBookRequest
	// var googleBookError request.GoogleBookErrorRequest
	err = decoder.Decode(&googleBook)
	if err != nil {
		panic(err)
	}

	return googleBook
}
