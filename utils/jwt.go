package utils

import (
	"time",
	"github.com/dgrijalva/jwt-go"
)

var jwt = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims {
		Username : username,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expirationTime.Unix(),
		},
	}
}