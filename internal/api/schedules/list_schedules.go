package schedules

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dagulv/gym-app/internal/model"
	"github.com/dagulv/gym-app/internal/db"
	// "github.com/dagulv/gym-app/internal/utils"
)

type schedule struct {
	Id uint8 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Color string `json:"color"`
}

func List(c echo.Context) error {
	schedules := make([]schedule, 0, 10)
	user := model.GetCurrent(c)

	resp, err := db.DBConn.Query("SELECT id, title, description, color FROM schedule WHERE userId = ?", user.Id)
	defer resp.Close()

	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"querymessage": err,
		})
	}

	for resp.Next() {
		schedule := schedule{}
		// var queryTimeCreated utils.RawTime 
		// var queryTimeUpdated utils.RawTime
		// &queryTimeCreated, &queryTimeUpdated
		// &schedule.Monday, &schedule.Tuesday, &schedule.Wednesday, &schedule.Thursday, &schedule.Friday, &schedule.Saturday, &schedule.Sunday
		if err = resp.Scan(&schedule.Id, &schedule.Title, &schedule.Description, &schedule.Color); err != nil {
			return c.JSON(http.StatusOK, echo.Map{
				"scanmessage": err,
			})
		}

		schedules = append(schedules, schedule)
	}
	return c.JSON(http.StatusOK, schedules)
	// return c.JSON(http.StatusOK, echo.Map{
	// 	"message": "hello from the echo server",
	// })

}
