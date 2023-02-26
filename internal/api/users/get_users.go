package users

import (
	"github.com/labstack/echo/v4"
	"github.com/dagulv/gym-app/internal/model"
)

func Get(c echo.Context) (err error) {
	user := model.UserContext{}
	id := c.Param("id")

	if id == "me" {
		currentUser := model.GetCurrent(c)
		return c.JSON(200, currentUser)
	} else {

	}
	return c.JSON(200, user)
}