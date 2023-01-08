package service

import "github.com/wachayathorn/golang-basic-structure/pkg/service/book"

type Service struct {
	BookService book.BookInterfaces
}

func InitService() *Service {
	return &Service{
		BookService: book.InitBookService(),
	}
}
