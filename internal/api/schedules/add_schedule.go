package schedules

import (
	"time"

	"github.com/labstack/echo/v4"
	// "github.com/dagulv/gym-app/internal/model"
	"github.com/dagulv/gym-app/internal/db"
)

func Create(c echo.Context) (err error) {

	_, err = db.DBConn.Exec("INSERT INTO schedule (userId, title, description, timeCreated, timeUpdated, color) VALUES (?, ?, ?, ?, ?, ?)", 1, "title", "desc", time.Now(), time.Now(), "ffffff")
	if err != nil {
        return err
    }

    return err
}