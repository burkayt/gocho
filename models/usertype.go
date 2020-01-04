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

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(user *User) (*User, error)
	DeleteUser(id int) error
	Login(user *User) (string, error)
}

type UserDao interface {
	User(id int) (*User, error)
	Users() (users []*User, err error)
	CreateUser(user *User) (*User, error)
	DeleteUser(id int) (err error)
	FindUserByEmail(email string) (*User, error)
}
