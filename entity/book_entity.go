package entity

import "time"

type Book struct {
	ID             	int
	UserID 			int
	Name 			string
	Genre 			string
	Release        	int64
	IsRead			bool
	CreatedAt      	time.Time
	UpdatedAt	   	time.Time
}