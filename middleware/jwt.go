package middleware

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var secretKey = []byte(os.Getenv("GRAM_TOKEN_SECRET"))

var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
