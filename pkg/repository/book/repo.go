package bookrepository

type Book struct {
	ID          uint `gorm:"primaryKey;index"`
	Title       string
	Author      string
	Description string
}
