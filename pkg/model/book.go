package model

type Book struct {
	Title       string `json:"title" example:"Title"`
	Author      string `json:"author" example:"Mr. John Corner"`
	Description string `json:"description" example:"This book for my family"`
}
