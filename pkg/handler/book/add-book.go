package bookhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wachayathorn/golang-basic-structure/pkg/model"
)

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
