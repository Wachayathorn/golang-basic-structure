package bookrepository

import "github.com/wachayathorn/golang-basic-structure/pkg/model"

const (
	CreateBook = `INSERT INTO book (title,author, description) VALUES ($1, $2, $3)`
)

func (b *BookRepositoryImpl) AddBook(book model.Book) (Book, error) {
	var result Book
	tx := b.Postgres.Begin()
	tx.Raw(CreateBook, book.Title, book.Author, book.Description)
	if err := tx.Commit(); err != nil {
		return result, err.Error
	}
	if err := tx.Row().Scan(&result); err != nil {
		return result, err
	}
	return result, nil
}
