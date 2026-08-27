package model

type User struct {
	Id       int     `gorm:"type:int;primary_key"`
	Name     string  `gorm:"type:varchar(150);not null"`
	Email    string  `gorm:"type:varchar(150);unique;not null"`
	Password string  `gorm:"type:varchar(60);not null"`
	Books    []*Book `gorm:"many2many:books_users;"`
}
