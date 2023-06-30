package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func CreateToken(username string) (string, error) {
	if username == "" {
		log.Fatal("Username tidak valid")
	}

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
