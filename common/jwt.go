package common

import (
	"time"

	"github.com/Wh0rigin/GraduationProject/bean"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("Wh0rigin")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user *bean.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Wh0rigin",
			Subject:   "AccessToken",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, err
}
