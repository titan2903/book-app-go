package repository

import (
	"book-app/entity"

	"gorm.io/gorm"
)

type RepositoryUser interface { //! public Interface
	AddUser(user entity.User) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	FindByID(ID int) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
}

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db}
}

func(r *repositoryUser) AddUser(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryUser) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryUser) FindByID(ID int) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryUser) Update(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func(r *repositoryUser) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}