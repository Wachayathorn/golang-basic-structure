package bookservice

import (
	bookrepository "github.com/wachayathorn/golang-basic-structure/pkg/repository/book"
)

func (b *BookService) GetBooks() []bookrepository.Book {
	return b.repo.GetBooks()
}
