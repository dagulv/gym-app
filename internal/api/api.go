package api

import (
	"github.com/labstack/echo/v4"

	"github.com/dagulv/gym-app/internal/api/schedules"
)

func Routes(api *echo.Group) {
	api.GET("/schedules", schedules.List)
	api.POST("/schedules", schedules.Create)
}