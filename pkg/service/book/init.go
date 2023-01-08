package book

import "github.com/wachayathorn/golang-basic-structure/pkg/model"

type BookInterfaces interface {
	AddBook() error
	GetBooks() []model.Book
}

type BookService struct {
}

func InitBookService() BookInterfaces {
	return &BookService{}
}
