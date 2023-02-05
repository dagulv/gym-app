package schedules

import (
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/dagulv/gym-app/internal/model"
	"github.com/dagulv/gym-app/internal/db"
)

func List(c echo.Context) (err error) {
	// schedules := make([]model.Schedule, 0, 10)
	err2 := db.DBConn.QueryRow("SELECT VERSION()")
	// defer resp.Close()

	if err2 != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": err2,
		})
	}

	// for resp.Next() {
	// 	var schedule model.Schedule
	// 	// &schedule.Monday, &schedule.Tuesday, &schedule.Wednesday, &schedule.Thursday, &schedule.Friday, &schedule.Saturday, &schedule.Sunday
	// 	err = resp.Scan(&schedule.Id, &schedule.UserId, &schedule.Title, &schedule.Description, &schedule.TimeCreated, &schedule.TimeUpdated, &schedule.Color)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	schedules = append(schedules, schedule)
	// }
	// return c.JSON(http.StatusOK, schedules)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "hello from the echo server",
	})

}