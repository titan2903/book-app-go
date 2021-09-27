package repository

import (
	"book-app/entity"
	"book-app/transport"
	"errors"

	"gorm.io/gorm"
)

type RepositoryBook interface {
	AddBook(book entity.Book) (entity.Book, error)
	FindAll(userID int, filterBook transport.FilterBook, limit, page int) ([]entity.Book, int64, error)
	FindByID(ID int) (entity.Book, error)
	UpdateBook(book entity.Book) (entity.Book, error)
	DeleteBook(ID int) (entity.Book, error)
}

type repositoryBook struct {
	db *gorm.DB
}

func NewRepositoryBook (db *gorm.DB) *repositoryBook {
	return &repositoryBook{db}
}

func (r *repositoryBook) AddBook(book entity.Book) (entity.Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repositoryBook) FindAll(userID int, filterBook transport.FilterBook, limit, page int) ([]entity.Book, int64, error) {
	var books []entity.Book

	var count int64

	offset := (page - 1) * limit

	query := r.db.Table("books").Count(&count)

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}

	if filterBook.Genre != "" {
		query = query.Where("genre LIKE ?", "%"+filterBook.Genre+"%")
	}

	if filterBook.StartYear != 0 && filterBook.EndYear != 0 {
		query = query
	}

	query = query.Count(&count).Find(&books)
	
	query = query.Limit(limit).Offset(offset).Find(&books)

	if query.Error != nil {
		return nil, 0, errors.New("Error Get Books from Database")
	}

	return books, count, nil
}

func (r *repositoryBook) FindByID(ID int) (entity.Book, error) {
	var book entity.Book

	err := r.db.Where("id = ?", ID).Find(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repositoryBook) UpdateBook(book entity.Book) (entity.Book, error) {
	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repositoryBook) DeleteBook(ID int) (entity.Book, error) {
	var book entity.Book

	err := r.db.Where("id = ?", ID).Delete(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

