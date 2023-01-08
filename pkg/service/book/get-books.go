package book

import "github.com/wachayathorn/golang-basic-structure/pkg/model"

func (b *BookService) GetBooks() []model.Book {
	return []model.Book{
		{
			Title:       "Title",
			Author:      "Author",
			Description: "Description",
		},
	}
}
