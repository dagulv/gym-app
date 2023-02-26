package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v4"
	"github.com/dagulv/gym-app/internal/api/schedules"
	"github.com/dagulv/gym-app/internal/api/auth"
	"github.com/dagulv/gym-app/internal/api/users"
	"github.com/dagulv/gym-app/internal/model"
	echojwt "github.com/labstack/echo-jwt/v4"
)

func Routes(api *echo.Group) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		TokenLookup: "header:Authorization:Bearer ,cookie:refreshToken",
		SigningKey: []byte("secret"),
	}

	api.POST("/login", auth.Login)
	api.POST("/signup", auth.Signup)

	api.GET("/users/:id", users.Get, echojwt.WithConfig(config))

	api.GET("/schedules", schedules.List, echojwt.WithConfig(config))
	api.POST("/schedules", schedules.Create, echojwt.WithConfig(config))
}

// func authentication(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(*auth.JwtCustomClaims)
// 		userContext := model.UserContext{}


// 		name := claims.Name
// 		return next(c)
// 	}
// }

func authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusOK, echo.Map{
				"token": ok,
			})
		}
		claims := token.Claims.(*model.JwtCustomClaims)

		user := model.UserContext{}

		user.Id = claims.Id
		user.Name = claims.Name
		c.Set("user", &user)

		return next(c)
	}
}