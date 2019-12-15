package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       int    `db:"ID"`
	Name     string `json:"name" db:"NAME"`
	Surname  string `json:"surname" db:"SURNAME"`
	Password string `json:"password" db:"PASSWORD"`
	Email    string `json:"email" db:"EMAIL"`
}

type JwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
