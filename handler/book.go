package book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-exercise/book"
	"net/http"
)

type Handler struct {
	bookService book.Service
}

func NewBookHandler(service book.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) BookGetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Dedi Fardiyanto",
		"bio":  "Fullstack Developer",
	})
}

func (h *Handler) BookHandlerPost(c *gin.Context) {
	var bookReq book.BookRequest

	err := c.ShouldBindJSON(&bookReq)
	if err != nil {
		var errorMessage []string

		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("error on field: %s should: %s", e.Field(), e.ActualTag())
			errorMessage = append(errorMessage, errorMsg)

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessage,
		})
		return
	}

	bookRes, err := h.bookService.Create(bookReq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":    true,
			"errorMsg": err,
			"data":     "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":    false,
		"errorMsg": "",
		"data":     bookRes,
	})

}
