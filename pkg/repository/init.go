package repository

import (
	bookrepository "github.com/wachayathorn/golang-basic-structure/pkg/repository/book"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	BookRepository bookrepository.BookRepository
}

func Init(postgres *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		BookRepository: bookrepository.Init(postgres),
	}
}
