package web

import (
	"github.com/labstack/echo"
	"log"
)

func RegisterHandlers() {
	e := echo.New()

	userHandler := CreateUserHandler()

	e.POST("/login", userHandler.loginHandler)

	e.GET("/user", userHandler.getAllUserHandler)
	e.GET("/user/:id", userHandler.getUserHandler)
	e.POST("/user", userHandler.postUserHandler)
	e.DELETE("/user/:id", userHandler.deleteUserHandler)

	log.Fatal(e.Start(":8080"))
}
