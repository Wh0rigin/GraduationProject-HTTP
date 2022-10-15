package po

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name       string //`gorm:"type:varchar(20);not null"`
	Isbn       string //`gorm:"type:varchar(11);not null"`
	Number     uint   //`gorm:"type:Integer;not null"`
	RentNumber uint   //`gorm:type:Integer;not null`
	Version    int    //`gorm:"type:Integer;not null"`
}

//constructor
func NewBookReference(name string, isbn string, number uint, rentNumber uint) *Book {
	return &Book{
		Name:       name,
		Isbn:       isbn,
		Number:     number,
		RentNumber: rentNumber,
		Version:    1,
	}
}

func NewBook(name string, isbn string, number uint, rentNumber uint) Book {
	return Book{
		Name:       name,
		Isbn:       isbn,
		Number:     number,
		RentNumber: rentNumber,
		Version:    1,
	}
}
