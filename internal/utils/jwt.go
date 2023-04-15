package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretkey = os.Getenv("SECRET_KEY")

func GenerateJWT(issuer int64) (token string, err error) {
	userid := strconv.Itoa(int(issuer))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		Issuer:    userid,
	})

	if secretkey == "" {
		secretkey = "secret"
	}

	return claims.SignedString([]byte(secretkey))
}

func ParseJWT(cookie string) (string, error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if secretkey == "" {
			secretkey = "secret"
		}
		return []byte(secretkey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
