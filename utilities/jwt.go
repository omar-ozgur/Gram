package utilities

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

var secretKey = []byte(os.Getenv("GRAM_TOKEN_SECRET"))

func GetClaims(tokenString string) map[string]interface{} {
	if tokenString == "" {
		return nil
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	return claims
}
