package repository

import (
	"book-app/entity"
	"book-app/transport"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type RepositoryBook interface {
	AddBook(book entity.Book) (entity.Book, error)
	FindAll(filterBook transport.FilterBook, limit, page int) ([]entity.Book, int64, error)
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

func (r *repositoryBook) FindAll(filterBook transport.FilterBook, limit, page int) ([]entity.Book, int64, error) {
	var books []entity.Book

	var count int64
	var join string

	isRead, _ := strconv.ParseBool(filterBook.IsRead)

	splitedTime := strings.Split(filterBook.Release, ",")

	if len(splitedTime) > 2 {
		return nil, 0, errors.New("Input date time max 2 input")
	}

	if len(splitedTime) == 1 {
		splited := strings.Split(filterBook.Release, "-")
		convertYear, _ := strconv.Atoi(splited[0])
		sumYear := convertYear + 1
		splited[0] = strconv.Itoa(sumYear)
		join = strings.Join(splited, "-")
	}

	offset := (page - 1) * limit

	query := r.db.Table("books").Count(&count)

	if filterBook.UserID != 0 {
		query = query.Where("user_id = ?", filterBook.UserID)
	}

	// fmt.Println("splited time: ", splitedTime)
	// fmt.Println("Release: ", filterBook.Release)

	// t, err := time.Parse(time.RFC3339, filterBook.Release)
	// fmt.Println("T: ",t.AddDate(1, 0, 0))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, 0, err
	// }

	if filterBook.Release != "" {
		if len(splitedTime) > 1 {
			query = query.Where("`release` BETWEEN ? AND ?", splitedTime[0], splitedTime[1])
		} else {
			fmt.Println("masuk sini")
			query = query.Where("`release` BETWEEN ? AND ?", filterBook.Release, join)
		}
	}

	if filterBook.Genre != "" {
		query = query.Where("genre LIKE ?", "%"+filterBook.Genre+"%")
	}

	if filterBook.IsRead != "" {
		query = query.Where("is_read = ?", isRead)
	}

	query = query.Count(&count).Find(&books)
	
	query = query.Limit(limit).Offset(offset).Find(&books)

	if query.Error != nil {
		return nil, 0, query.Error
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

