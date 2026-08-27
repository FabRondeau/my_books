package model

type Author struct {
	Id       int     `gorm:"type:int;primary_key"`
	FullName string  `gorm:"type:varchar(150);not null"`
	Books    []*Book `gorm:"many2many:books_authors;"`
}
