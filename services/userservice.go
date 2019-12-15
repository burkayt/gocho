package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
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
	//TODO bcrypt password
	return dao.CreateUser(user)

}

func DeleteUser(id int) error {
	return dao.DeleteUser(id)
}

func Login(user *models.User) (string, error) {
	secret := viper.GetString("jwt.secret")

	foundUser, err := dao.FindUserByEmail(user.Email)

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
