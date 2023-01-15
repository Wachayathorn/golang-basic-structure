package bookservice

import (
	"github.com/wachayathorn/golang-basic-structure/pkg/model"
	bookrepository "github.com/wachayathorn/golang-basic-structure/pkg/repository/book"
)

type BookInterfaces interface {
	AddBook(book model.Book) (bookrepository.Book, error)
	GetBooks() []bookrepository.Book
}

type BookService struct {
	repo bookrepository.BookRepository
}

func InitBookService(repo bookrepository.BookRepository) BookInterfaces {
	return &BookService{
		repo: repo,
	}
}
