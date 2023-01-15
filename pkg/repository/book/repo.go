package bookrepository

type Book struct {
	ID          int    `gorm:"column:id"`
	Title       string `gorm:"column:title"`
	Author      string `gorm:"column:author"`
	Description string `gorm:"column:description"`
}
