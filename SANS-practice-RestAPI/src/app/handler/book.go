package handler

import (
	"fmt"
	"net/http"
	"sans-api-pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(b book.Service) *bookHandler {
	return &bookHandler{b}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "name",
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := h.bookService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":     false,
			"status": http.StatusInternalServerError,
			"errors": err,
		})

		return
	}

	bookResponse := convBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"status": http.StatusOK,
		"data":   bookResponse,
	})

}

func (h *bookHandler) GetBooks(c *gin.Context) {

	s := c.Query("s")
	if s != "" {
		bookFindLike, _ := h.bookService.FindLIKE(s)
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"status": http.StatusOK,
			"data":   bookFindLike,
		})
		return
	}
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":     false,
			"status": http.StatusInternalServerError,
			"errors": err,
		})

		return
	}

	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"status": http.StatusOK,
		"data":   booksResponse,
	})

}

func (h *bookHandler) CreateBook(c *gin.Context) {

	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	tmpErrors := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("%s is %s", e.Field(), e.ActualTag())
			tmpErrors = append(tmpErrors, message)
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":     false,
			"status": http.StatusInternalServerError,
			"errors": tmpErrors,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"status": http.StatusOK,
		"data":   book,
	})
}
func (h *bookHandler) UpdateBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	tmpErrors := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("%s is %s", e.Field(), e.ActualTag())
			tmpErrors = append(tmpErrors, message)
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":     false,
			"status": http.StatusInternalServerError,
			"errors": tmpErrors,
		})
		return
	}

	book, err := h.bookService.Update(bookRequest, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"status": http.StatusOK,
		"data":   book,
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	bookDeleted, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"status": http.StatusOK,
		"data":   bookDeleted,
	})

}

func convBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:        b.ID,
		Title:     b.Title,
		Price:     b.Price,
		Rating:    b.Rating,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}

}
