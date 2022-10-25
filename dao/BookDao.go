package dao

import (
	"errors"

	"github.com/Wh0rigin/GraduationProject/common"
	"github.com/Wh0rigin/GraduationProject/po"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *po.Book) {
	db.Create(book)
}

func DeletetBook(db *gorm.DB, isbn string) error {
	book, err := GetBookByIsbn(db, isbn)
	if err != nil {
		return err
	}
	db.Delete(book)
	return err
}

func IsIsbnExist(db *gorm.DB, isbn string) bool {
	var book po.Book
	db.Where("isbn = ?", isbn).First(&book)
	if book.ID != 0 {
		return true
	}
	return false
}

func GetBookByIsbn(db *gorm.DB, isbn string) (*po.Book, error) {
	var book po.Book
	db.Where("isbn=?", isbn).First(&book)
	if book.ID != 0 {
		return &book, nil
	}
	return nil, errors.New("dml:null value")
}

func GetBookByID(db *gorm.DB, id uint) (*po.Book, error) {
	var book po.Book
	db.Where("ID=?", id).First(&book)
	if book.ID != 0 {
		return &book, nil
	}
	return nil, errors.New("dml:null value")
}

func UpdateBook(db *gorm.DB, book *po.Book) error {
	nowBook, err := GetBookByID(common.GetDb(), book.ID)
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

func GetAllBook(db *gorm.DB) (books []po.Book) {
	db.Find(&books)
	return books
}

// select book by hazy isbn
func SeleteBook(db *gorm.DB, isbn string) (books []po.Book) {
	db.Where("isbn like ?", isbn+"%").Find(&books)
	return books
}
