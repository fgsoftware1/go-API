package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Writer string  `json:"writer"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Blue Train", Writer: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Writer: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Clifford Brown", Writer: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("api/books", getBooks)
	router.GET("api/books/:id", getBookByID)
	router.POST("api/books", postBooks)
	router.PATCH("api/books/:id", postBooks)

	router.Run("localhost:8080")
}

//GET api/books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

//POST api/books
func postBooks(c *gin.Context) {
	var newBook book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

//GET api/books/:id
func getBookByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range books {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
}

//PATCH api/books/:id
func UpdateBook(c *gin.Context) {
	var altBook book

	if err := c.ShouldBindJSON(&altBook); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated!"})
}

