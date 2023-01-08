package service

import "github.com/wachayathorn/golang-basic-structure/pkg/service/book"

type Service struct {
	BookService book.BookInterfaces
}

func Init() *Service {
	return &Service{
		BookService: book.Init(),
	}
}
