package web

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"gocho/models"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService models.UserService
}

func NewUserHandler(service models.UserService) *UserHandler {
	return &UserHandler{service}
}

func (userHandler *UserHandler) deleteUserHandler(c echo.Context) error {
	param := c.Param("id")

	userId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = userHandler.userService.DeleteUser(userId)

	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, nil)

}

func (userHandler *UserHandler) getUserHandler(c echo.Context) error {
	param := c.Param("id")

	userId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user, err := userHandler.userService.User(userId)

	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (userHandler *UserHandler) getAllUserHandler(c echo.Context) error {
	users, err := userHandler.userService.Users()

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, users)
}

func (userHandler *UserHandler) postUserHandler(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err := userHandler.userService.CreateUser(&user)

	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (userHandler *UserHandler) loginHandler(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if len(user.Password) <= 0 || len(user.Email) <= 0 {
		return errors.New("Email and Password is mandatory")
	}

	token, err := userHandler.userService.Login(&user)

	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, token)
}
