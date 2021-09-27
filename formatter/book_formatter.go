package formatter

import "book-app/entity"

type BookFormatter struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Genre string `json:"genre"`
	Release int64 `json:"release"`
}

func FormatBook(book entity.Book) BookFormatter {
	formatter := BookFormatter{
		ID:              	book.ID,
		UserID:           	book.UserID,
		Name:             	book.Name,
		Genre: 				book.Genre,
		Release:       		book.Release,
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
