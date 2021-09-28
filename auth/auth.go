package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)


type AuthService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {

}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func NewService() *jwtService { //! bisa memanggil generate token dari package mana pun
	return &jwtService{}
}

func(s *jwtService) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID //! data yang ingin di masukkan ke token
	payload["exp"] = time.Now().Add(time.Minute * 120).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //! generate token

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func(s *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodeToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	}) //! untuk melakukan validasi perlu parse telebih dahulu tokennya

	if err != nil {
		return token, err
	}

	return token, nil //! berhasil di validasi
}