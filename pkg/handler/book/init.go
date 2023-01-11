package book

import (
	"github.com/gin-gonic/gin"
	"github.com/wachayathorn/golang-basic-structure/pkg/handler/book/httpinfo"
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
	book := r.Group(httpinfo.GroupV1)
	{
		book.POST(httpinfo.AddBook, b.AddBook, middleware.Middleware)
		book.GET(httpinfo.GetBooks, b.GetBooks, middleware.Middleware)
	}
}
