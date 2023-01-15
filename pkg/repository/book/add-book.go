package bookrepository

import (
	"github.com/wachayathorn/golang-basic-structure/pkg/model"
)

const (
	CreateBook = `INSERT INTO book (title,author,description) VALUES (?, ?, ?) RETURNING *`
)

func (b *BookRepositoryImpl) AddBook(book model.Book) (Book, error) {
	var result Book
	tx := b.Postgres.Begin()
	insert := tx.Raw(CreateBook, book.Title, book.Author, book.Description)
	if err := insert.Error; err != nil {
		tx.Rollback()
		return result, err
	}
	if err := insert.Scan(&result).Error; err != nil {
		tx.Rollback()
		return result, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return result, err
	}
	return result, nil
}
