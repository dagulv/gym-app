package main

import(
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/dagulv/gym-app/internal/exercise"
)
func main() {
	r := echo.New()
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	rg := r.Group("")
	exercise.RegisterHandlers(rg)
	// app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
    //     Root: "/frontend/build",
	// 	Browse: true,
    // }))

	log.Fatal(r.Start(":8080"))
}
