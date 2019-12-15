package web

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"gocho/models"
	"gocho/services"
	"net/http"
	"strconv"
)

func deleteUserHandler(c echo.Context) error {
	param := c.Param("id")

	userId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = services.DeleteUser(userId)

	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, nil)

}

func getUserHandler(c echo.Context) error {
	param := c.Param("id")

	userId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user, err := services.User(userId)

	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func getAllUserHandler(c echo.Context) error {
	users, err := services.Users()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, users)
}

func postUserHandler(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := services.CreateUser(&user)

	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func loginHandler(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if len(user.Password) <= 0 || len(user.Email) <= 0 {
		return errors.New("Email and Password is mandatory")
	}

	token, err := services.Login(&user)

	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, token)
}
