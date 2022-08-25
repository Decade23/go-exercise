package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Dedi Fardiyanto",
			"bio":  "Fullstack Developer",
		})
	})

	router.POST("/books", bookHandler)

	router.Run(":80")
}

func bookHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		//fmt.Println(err)
		//c.JSON(http.StatusBadRequest, err)
		//return
		errorMessage := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("error on field: %s should: %s", e.Field(), e.ActualTag())
			fmt.Println(errorMsg)
			errorMessage := append(errorMessage, errorMsg)

			c.JSON(http.StatusBadRequest, errorMessage)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})

}

type BookInput struct {
	Title string          `json:"title" binding:"required"`
	Price jsoniter.Number `json:"price" binding:"required,number"`
}
