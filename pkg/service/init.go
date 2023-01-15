package service

import (
	"github.com/wachayathorn/golang-basic-structure/pkg/repository"
	bookservice "github.com/wachayathorn/golang-basic-structure/pkg/service/book"
)

type Service struct {
	BookService bookservice.BookInterfaces
}

func Init(repo *repository.RepositoryImpl) *Service {
	return &Service{
		BookService: bookservice.InitBookService(repo.BookRepository),
	}
}
