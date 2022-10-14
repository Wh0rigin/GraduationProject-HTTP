package dao

import (
	"github.com/Wh0rigin/GraduationProject/bean"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *bean.Book) {
	db.Create(book)
}

func IsIsbnExist(db *gorm.DB, isbn string) bool {
	var book bean.Book
	db.Where("isbn = ?", isbn).First(&book)
	if book.ID != 0 {
		return true
	}
	return false
}
