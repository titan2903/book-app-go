package entity

import "time"

type User struct {
	ID             	int
	Name 			string
	Username 		string
	Password 		string
	CreatedAt      	time.Time
	UpdatedAt	   	time.Time
}