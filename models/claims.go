package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserEmail string `json:"user_email""`
	jwt.StandardClaims
}
