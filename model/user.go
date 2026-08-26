package model

type User struct {
	Id       int     `gorm:"type:int;primary_key"`
	Name     string  `gorm:"type:varchar(150);not null"`
	Email    string  `gorm:"type:varchar(150);unique;not null"`
	Password string  `gorm:"type:varchar(60);not null"`
	Books    []*Book `gorm:"many2many:books_users;"`
}
type Publisher struct {
	Id    int    `gorm:"type:int;primary_key"`
	Name  string `gorm:"type:varchar(150);not null"`
	Books []Book `gorm:"foreignkey:PublisherID"`
}

type Author struct {
	Id       int     `gorm:"type:int;primary_key"`
	FullName string  `gorm:"type:varchar(150);not null"`
	Books    []*Book `gorm:"many2many:books_authors;"`
}
type Book struct {
	Id          int    `gorm:"type:int;primary_key"`
	ISBN13      string `gorm:"type:varchar(13);not null"`
	Title       string `gorm:"type:varchar(150);not null"`
	Subtitle    string `gorm:"type:varchar(200);not null"`
	PublisherID int
	Authors     []*Author `gorm:"many2many:books_authors;"`
	Users       []*User   `gorm:"many2many:books_users;"`
}
