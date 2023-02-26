package schedules

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/dagulv/gym-app/internal/model"
	"github.com/dagulv/gym-app/internal/db"
)

func Create(c echo.Context) (err error) {

	schedule := model.Schedule{}
	err = c.Bind(&schedule); 
	
	if err != nil {
		return err
	}

	_, err = db.DBConn.Exec("INSERT INTO schedule (userId, title, description, color, timeCreated, timeUpdated) VALUES (?, ?, ?, ?, ?, ?)", schedule.UserId, schedule.Title, schedule.Description, "ffffff", time.Now(), time.Now())
	if err != nil {
        return err
    }

    return nil
}