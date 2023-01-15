package bookrepository

import (
	"github.com/wachayathorn/golang-basic-structure/pkg/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	AddBook(book model.Book) (Book, error)
	GetBooks() []Book
}

type BookRepositoryImpl struct {
	Postgres *gorm.DB
}

func Init(postgres *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		Postgres: postgres,
	}
}
