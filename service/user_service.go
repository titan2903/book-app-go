package service

import (
	"book-app/entity"
	"book-app/repository"
	"book-app/transport"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


type ServiceUser interface { //! bisnis logic
	RegisterUser(input transport.RegisterUserInput) (entity.User, error)
	Login(input transport.LoginInput) (entity.User, error)
	IsUserNameAvailable(input transport.RegisterUserInput) (bool, error)
	GetUserByID(ID int) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(input transport.FormUpdateUserInput) (entity.User, error)
}

type serviceUser struct {
	repository repository.RepositoryUser
}

func NewServiceUser(repository repository.RepositoryUser) *serviceUser {
	return &serviceUser{repository}
}

func(s *serviceUser) RegisterUser(input transport.RegisterUserInput) (entity.User, error) {
	fmt.Printf("Username user: %s",input.Username)
	userData, _ := s.repository.FindByUsername(input.Username)

	fmt.Printf("User data: %+v",userData)
	if userData.ID != 0 { //! jika ada user yang menggunakan username yang sama
		return userData, errors.New("Username has been taken")
	}
	
	user := entity.User{}
	user.Name = input.Name
	user.Username = input.Username
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.AddUser(user)

	if err != nil {
		return newUser ,err
	}

	return newUser, nil
}

func (s *serviceUser) Login(input transport.LoginInput) (entity.User, error) {
	username := input.Username
	password := input.Password

	//! checking username user in database
	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	//! compare password in database with input password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *serviceUser) IsUserNameAvailable(input transport.RegisterUserInput) (bool, error) {
	username := input.Username

	//! checking jika username ada di database
	user, _ := s.repository.FindByUsername(username)

	if user.ID == 0 { //! jika tidak ada user
		return true, nil
	}

	return false, nil
}

func (s *serviceUser) GetUserByID(ID int) (entity.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with that ID")
	}

	return user, nil
}

func (s *serviceUser) GetAllUsers() ([]entity.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *serviceUser) UpdateUser(input transport.FormUpdateUserInput) (entity.User, error) {
	user, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	if user.ID != input.User.ID {
		return user, errors.New("not owner of book")
	}

	user.Name = input.Name
	user.Username = input.Username

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}