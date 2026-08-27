package model

type Book struct {
	Id          int    `gorm:"type:int;primary_key"`
	ISBN13      string `gorm:"type:varchar(13);not null"`
	Title       string `gorm:"type:varchar(150);not null"`
	Subtitle    string `gorm:"type:varchar(200);not null"`
	PublisherID int
	Authors     []*Author `gorm:"many2many:books_authors;"`
	Users       []*User   `gorm:"many2many:books_users;"`
}
