package model

import "github.com/golang-jwt/jwt/v4"

type User struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
}

type CustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
