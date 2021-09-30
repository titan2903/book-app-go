package formatter

import (
	"book-app/entity"
	"time"
)

type BookFormatter struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Genre string `json:"genre"`
	Release time.Time `json:"release"`
	IsRead bool `json:"is_read"`
}

func FormatBook(book entity.Book) BookFormatter {
	formatter := BookFormatter{
		ID:              	book.ID,
		UserID:           	book.UserID,
		Name:             	book.Name,
		Genre: 				book.Genre,
		Release:       		book.Release,
		IsRead: 			book.IsRead,
	}

	return formatter
}


func FormatBooks(books []entity.Book) []BookFormatter { //! mengembalikan array of object
	booksFormatter := []BookFormatter{}

	for _, book := range books {
		bookFormatter := FormatBook(book)
		booksFormatter = append(booksFormatter, bookFormatter)
	}

	return booksFormatter
}

func FormatBookDetail(book entity.Book) BookFormatter {
	bookDetailFormatter := BookFormatter{}
	bookDetailFormatter.ID = book.ID
	bookDetailFormatter.Name = book.Name
	bookDetailFormatter.Genre = book.Genre
	bookDetailFormatter.Release = book.Release
	bookDetailFormatter.UserID = book.UserID

	return bookDetailFormatter
}
