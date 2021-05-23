package db

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId int64 `json:"id"`
	jwt.StandardClaims
}
