package entity

import "time"

type Book struct {
	ID             	int
	UserID 			int
	Name 			string
	Genre 			string
	Release        	int64
	CreatedAt      	time.Time
	UpdatedAt	   	time.Time
}