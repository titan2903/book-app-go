package formatter

import "book-app/entity"

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Username string `json:"username"`
	Token      string `json:"token"`
}

type UserFetchFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Username string `json:"username"`
}

func FormatUser(user entity.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Username: 	user.Username,
		Token:      token,
	}

	return formatter
}

func FormatFetchUser(user entity.User) UserFetchFormatter {
	formatter := UserFetchFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Username: 	user.Username,
	}

	return formatter
}