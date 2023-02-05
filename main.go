package main

import(
	"log"
	"net/http"
	// "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/dagulv/gym-app/internal/db"
	"github.com/dagulv/gym-app/internal/server"
	"github.com/dagulv/gym-app/internal/api"
	
)
func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	
	err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	s := server.Connect()

	api.Routes(s.Group(
		"/api",
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
			AllowCredentials: true,
		}),
	))

	log.Fatal(server.Start(s))
}
