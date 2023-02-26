package main

import(
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	// echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/dagulv/gym-app/internal/db"
	// "github.com/dagulv/gym-app/internal/server"
	"github.com/dagulv/gym-app/internal/api"
	// "github.com/dagulv/gym-app/internal/api/auth"
	
)
func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	
	err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	s := echo.New()

	g := s.Group("/api")

	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
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
		middleware.Decompress(),
		middleware.Gzip(),
	)

	
	// s.Use(middleware.CORS());
	// Configure middleware with the custom claims type
	// config := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(auth.JwtCustomClaims)
	// 	},
	// 	SigningKey: []byte("secret"),
	// }
	// g.Use(echojwt.WithConfig(config))

	api.Routes(g)

	log.Fatal(s.Start(":8080"))
}
