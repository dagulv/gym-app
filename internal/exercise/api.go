package exercise

import (
	"net/http"
	"github.com/dagulv/gym-app/internal/model"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(r *echo.Group) {
	
	r.GET("/schema", handleSchedules)
	r.GET("/hello.json", handleHello)
}

func handleSchedules(c echo.Context) error {
	data := model.Schedule {
		Id: 132,
	}
	// data := make(model.Schedule, 0, 10);
	// data = append(data, 
	// {
	// 	Id: 123,
	// 	Title: "Schema 1",
	// 	Description: "Beskrivning",
	// 	Monday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 		"exercise-4": "Övning 4",
	// 	},
	// 	Tuesday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 		"exercise-4": "Övning 4",
	// 		"exercise-5": "Övning 5",
	// 	},
	// 	Wednesday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 		"exercise-4": "Övning 4",
	// 		"exercise-5": "Övning 5",
	// 		"exercise-6": "Övning 6",
	// 	},
	// 	Thursday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 		"exercise-4": "Övning 4",
	// 		"exercise-5": "Övning 5",
	// 		"exercise-6": "Övning 6",
	// 		"exercise-7": "Övning 7",
	// 	},
	// 	Friday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 	},
	// 	Saturday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 	},
	// 	Sunday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 	},
	// });
	// {
	// 	Id: 123,
	// 	Title: "Schema 1",
	// 	Description: "Beskrivning",
	// 	Monday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 		"exercise-4": "Övning 4",
	// 	},
	// 	Tuesday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 		"exercise-4": "Övning 4",
	// 		"exercise-5": "Övning 5",
	// 	},
	// 	Wednesday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 		"exercise-4": "Övning 4",
	// 		"exercise-5": "Övning 5",
	// 		"exercise-6": "Övning 6",
	// 	},
	// 	Thursday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 		"exercise-4": "Övning 4",
	// 		"exercise-5": "Övning 5",
	// 		"exercise-6": "Övning 6",
	// 		"exercise-7": "Övning 7",
	// 	},
	// 	Friday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 	},
	// 	Saturday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 	},
	// 	Sunday: map[string]interface{}{
	// 		"exercise-1": "Övning 1",
	// 		"exercise-2": "Övning 2",
	// 		"exercise-3": "Övning 3",
	// 	},
	// }
	
	return c.JSON(http.StatusOK, data)
  }

func handleHello(c echo.Context) error  {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "hello from the echo server",
	})
}