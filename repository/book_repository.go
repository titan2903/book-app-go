package repository

import (
	"book-app/entity"

	"gorm.io/gorm"
)

type RepositoryBook interface {
	AddBook(book entity.Book) (entity.Book, error)
	FindAll(userID int, genre string) ([]entity.Book, error)
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

func (r *repositoryBook) FindAll(userID int, genre string) ([]entity.Book, error) {
	var books []entity.Book

	if userID != 0 {
		err := r.db.Where("user_id = ?", userID).Find(&books).Error
		if err != nil {
			return books, err
		}

		return books, nil
	} else if genre != "" {
		err := r.db.Where("genre = ?", genre).Find(&books).Error
		if err != nil {
			return books, err
		}

		return books, nil
	} else {
		err := r.db.Find(&books).Error
		if err != nil {
			return books, err
		}
	}

	

	return books, nil
}

func (r *repositoryBook) FindByUserID(userID int) ([]entity.Book, error) {
	var book []entity.Book

	

	return book, nil
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

