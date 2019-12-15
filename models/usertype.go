package models

type User struct {
	Id      int    `db:"ID"`
	Name    string `json:"name" db:"NAME"`
	Surname string `json:"surname" db:"SURNAME"`
}

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(user *User) (*User, error)
	DeleteUser(id int) error
}
