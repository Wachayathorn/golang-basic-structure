package bookservice

import (
	"github.com/wachayathorn/golang-basic-structure/pkg/model"
	bookrepository "github.com/wachayathorn/golang-basic-structure/pkg/repository/book"
)

func (b *BookService) AddBook(book model.Book) (bookrepository.Book, error) {
	result, err := b.repo.AddBook(book)
	if err != nil {
		return result, err
	}
	return result, nil
}
