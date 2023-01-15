package bookhandler

import (
	"github.com/gin-gonic/gin"
	bookhttpinfo "github.com/wachayathorn/golang-basic-structure/pkg/handler/book/httpinfo"
	"github.com/wachayathorn/golang-basic-structure/pkg/handler/middleware"
	"github.com/wachayathorn/golang-basic-structure/pkg/service"
	bookservice "github.com/wachayathorn/golang-basic-structure/pkg/service/book"
)

type BookHandler struct {
	BookService bookservice.BookInterfaces
}

func InitV1(r *gin.Engine, svc *service.Service) {
	b := &BookHandler{
		BookService: svc.BookService,
	}
	book := r.Group(bookhttpinfo.GroupV1)
	{
		book.POST(bookhttpinfo.AddBook, b.AddBook, middleware.Middleware)
		book.GET(bookhttpinfo.GetBooks, b.GetBooks, middleware.Middleware)
	}
}
