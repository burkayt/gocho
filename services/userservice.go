package services

import (
	"gocho/dao"
	"gocho/models"
)

func User(id int) (*models.User, error) {
	return dao.User(id)
}

func Users() ([]*models.User, error) {
	return dao.Users()
}

func CreateUser(user *models.User) (*models.User, error) {
	return dao.CreateUser(user)

}

func DeleteUser(id int) error {
	return dao.DeleteUser(id)
}
