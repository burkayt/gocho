package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gocho/models"
)

type UserService struct {
	userDao models.UserDao
}

func NewUserService(userDao models.UserDao) models.UserService {
	return &UserService{userDao}
}

func (u *UserService) User(id int) (*models.User, error) {
	return u.userDao.User(id)
}

func (u *UserService) Users() ([]*models.User, error) {
	return u.userDao.Users()
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	//TODO bcrypt password
	return u.userDao.CreateUser(user)

}

func (u *UserService) DeleteUser(id int) error {
	return u.userDao.DeleteUser(id)
}

func (u *UserService) Login(user *models.User) (string, error) {
	secret := viper.GetString("jwt.secret")

	foundUser, err := u.userDao.FindUserByEmail(user.Email)

	if err != nil {
		return "", err
	}

	if foundUser.Password != user.Password {
		return "", errors.New("password mismatch")
	}

	claim := models.JwtClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(secret))
}
