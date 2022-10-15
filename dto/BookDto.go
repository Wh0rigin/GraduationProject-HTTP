package dto

import "github.com/Wh0rigin/GraduationProject/bean"

type BookDto struct {
	Name            string `Json:"name"`
	Isbn            string `Json:"Isbn"`
	Number          uint   `Json:"number"`          //	allNumber
	AvailableNumber uint   `Json:"availableNumber"` //	availableNumber
}

func NewBookDto(book bean.Book) *BookDto {
	return &BookDto{
		Name:            book.Name,
		Isbn:            book.Isbn,
		Number:          book.Number,
		AvailableNumber: book.Number - book.RentNumber,
	}
}

func NewBookDtoByReference(book *bean.Book) *BookDto {
	return &BookDto{
		Name:            book.Name,
		Isbn:            book.Isbn,
		Number:          book.Number,
		AvailableNumber: book.Number - book.RentNumber,
	}
}

func NewBookDtosByArray(books []bean.Book) (bookDtos []BookDto) {
	bookDtos = make([]BookDto, len(books))
	for i := 0; i < len(books); i++ {
		bookDtos[i] = *NewBookDto(books[i])
	}
	return bookDtos
}
