package bookservice

import (
	log "github.com/sirupsen/logrus"
	"github.com/wachayathorn/golang-basic-structure/pkg/model"
	bookrepository "github.com/wachayathorn/golang-basic-structure/pkg/repository/book"
)

func (b *BookService) AddBook(book model.Book) (bookrepository.Book, error) {
	result, err := b.repo.AddBook(book)
	if err != nil {
		log.Errorf("service - AddBook : repo add book got error: %s", err.Error())
		return result, err
	}
	return result, nil
}
