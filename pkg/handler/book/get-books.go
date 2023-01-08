package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BookHandler) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, h.BookService.GetBooks())
}
