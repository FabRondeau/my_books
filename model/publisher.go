package model

type Publisher struct {
	Id    int    `gorm:"type:int;primary_key"`
	Name  string `gorm:"type:varchar(150);not null"`
	Books []Book `gorm:"foreignkey:PublisherID"`
}
