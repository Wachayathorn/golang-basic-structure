package bookhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wachayathorn/golang-basic-structure/pkg/model"
)

// AddBook godoc
// @summary Add Book
// @description Add Book
// @tags book
// @id AddBook
// @accept json
// @produce json
// @param Book body model.Book true "Book data to be created"
// @response 201 {object} bookrepository.Book "Created"
// @response 400 {string} string "Bad Request"
// @response 500 {string} string "Internal Server Error"
// @Router /api/v1/book [post]
func (h *BookHandler) AddBook(c *gin.Context) {
	body := model.Book{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.BookService.AddBook(body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
