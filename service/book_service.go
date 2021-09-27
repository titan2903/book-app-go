package service

import (
	"book-app/entity"
	"book-app/repository"
	"book-app/transport"
	"errors"
)

type ServiceBook interface {
	AddBook(input transport.InputDataBook) (entity.Book, error)
	GetBooks(UserID int, filterBook transport.FilterBook, limit, page int) ([]entity.Book, int64, error)
	FindByID(input int) (entity.Book, error)
	UpdateBook(inputID transport.InputDetailIdBook, inputData transport.InputDataBookUpdate) (entity.Book, error)
	DeleteBook(inputID transport.InputDetailIdBook, inputData transport.InputDetailIdBook) (entity.Book, error)
}

type serviceBook struct {
	repository repository.RepositoryBook
}

func NewBookService(repository repository.RepositoryBook) *serviceBook {
	return &serviceBook{repository}
}

func(s *serviceBook) AddBook(input transport.InputDataBook) (entity.Book, error) {
	book := entity.Book{}
	book.Name = input.Name
	book.Genre = input.Genre
	book.Release = input.Release
	book.UserID = input.User.ID
	book.IsRead = false

	newBook, err := s.repository.AddBook(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func(s *serviceBook) GetBooks(UserID int, filterBook transport.FilterBook, limit, page int) ([]entity.Book, int64, error) {

	books, count, err := s.repository.FindAll(UserID, filterBook, limit, page)
	if err != nil {
		return books, 0, err
	}

	return books, count, nil
}

func(s *serviceBook) FindByID(input int) (entity.Book, error) {
	book, err := s.repository.FindByID(input)

	if err != nil {
		return book, err
	}

	return book, nil
}

func(s *serviceBook) UpdateBook(inputID transport.InputDetailIdBook, inputData transport.InputDataBookUpdate) (entity.Book, error) {
	book, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return book, err
	}

	if book.UserID != inputData.User.ID {
		return book, errors.New("not owner of book")
	}

	if inputData.Name != "" {
		book.Name = inputData.Name
	}

	if inputData.Genre != "" {
		book.Genre = inputData.Genre
	}

	if inputData.Release != 0 {
		book.Release = inputData.Release
	}

	if inputData.IsRead || !inputData.IsRead {
		book.IsRead = inputData.IsRead
	}

	updated, err := s.repository.UpdateBook(book)
	if err != nil {
		return updated, err
	}

	return updated, nil
}

func(s *serviceBook) DeleteBook(inputID, inputData transport.InputDetailIdBook) (entity.Book, error) {
	book, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return book, err
	}

	if book.UserID != inputData.User.ID {
		return book, errors.New("not owner of book")
	}

	deleted, err := s.repository.DeleteBook(inputID.ID)
	if err != nil {
		return deleted, err
	}


	return deleted, nil
}