package server

import (
	"github.com/labstack/echo/v4"
)

func Connect() (server *echo.Echo) {
	return echo.New()
}

func Start(server *echo.Echo) (err error) {
	err = server.Start(":8080")
	return
}