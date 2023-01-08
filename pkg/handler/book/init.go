package book

import (
	"github.com/gin-gonic/gin"
	"github.com/wachayathorn/golang-basic-structure/pkg/handler/middleware"
	"github.com/wachayathorn/golang-basic-structure/pkg/service"
	"github.com/wachayathorn/golang-basic-structure/pkg/service/book"
)

type BookHandler struct {
	BookService book.BookInterfaces
}

func InitApiV1(r *gin.Engine, svc *service.Service) {
	b := &BookHandler{
		BookService: svc.BookService,
	}
	book := r.Group("/api/v1/book")
	{
		book.POST("", b.AddBook, middleware.Middleware)
		book.GET("", b.GetBooks, middleware.Middleware)
	}
}
