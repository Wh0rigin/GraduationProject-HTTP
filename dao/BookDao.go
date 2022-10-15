package dao

import (
	"errors"

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

func GetBookByIsbn(db *gorm.DB, isbn string) (*bean.Book, error) {
	var book bean.Book
	db.Where("isbn=?", isbn).First(&book)
	if book.ID != 0 {
		return &book, nil
	}
	return nil, errors.New("dml:null value")
}

func GetBookByID(db *gorm.DB, id uint) (*bean.Book, error) {
	var book bean.Book
	db.Where("ID=?", id).First(&book)
	if book.ID != 0 {
		return &book, nil
	}
	return nil, errors.New("dml:null value")
}

func UpdateBook(db *gorm.DB, book *bean.Book) error {
	nowBook, err := GetBookByID(GetDb(), book.ID)
	if err != nil {
		return err
	}
	// positive lock
	if nowBook.Version != book.Version {
		return errors.New("dml:version clash")
	}

	book.Version += 1
	db.Save(book)
	return nil
}
