package transport

import "book-app/entity"

//! struct yang digunakan untuk mapping dari inputan user
type RegisterUserInput struct { 
	Name       	string `json:"name" binding:"required"`
	Username 	string `json:"username" binding:"required"`
	Password   	string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password 	string `json:"password" form:"password" binding:"required"`
}

type CheckUsernameInput struct {
	Username string `json:"username" binding:"required"`
}

type FormCreateUserInput struct {
	Name       string `form:"name" binding:"required"`
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Error      error
}

type FormUpdateUserInput struct {
	ID         		int
	Name       		string `form:"name" binding:"required"`
	Username      	string `form:"username" binding:"required"`
	User entity.User
	Error      		error
}

type InputDataBook struct {
	Name string `json:"name" binding:"required"`
	Genre string `json:"genre"`
	Release int64 `json:"release"`
	User entity.User
}

type InputDetailIdBook struct {
	ID int `uri:"id" binding:"required"`
	User entity.User
}