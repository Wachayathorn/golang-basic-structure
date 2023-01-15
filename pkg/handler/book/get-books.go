package bookhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @summary Get books
// @description Get books
// @tags book
// @id GetBooks
// @accept json
// @produce json
// @response 200 {object} []bookrepository.Book "Created"
// @Router /api/v1/book [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, h.BookService.GetBooks())
}
