package web

import (
	"github.com/labstack/echo"
	"log"
)

func RegisterHandlers() {
	e := echo.New()

	e.POST("/login", loginHandler)

	e.GET("/user", getAllUserHandler)
	e.GET("/user/:id", getUserHandler)
	e.POST("/user", postUserHandler)
	e.DELETE("/user/:id", deleteUserHandler)

	log.Fatal(e.Start(":8080"))
}
