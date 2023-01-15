package bookrepository

const (
	GetBooks = `SELECT * FROM book ORDER BY book.id`
)

func (b *BookRepositoryImpl) GetBooks() []Book {
	var result []Book
	b.Postgres.Raw(GetBooks).Scan(&result)
	return result
}
