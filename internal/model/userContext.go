package model

import (
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Id int `json:"id"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

type UserContext struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func GetCurrent(c echo.Context) UserContext {
	token := c.Get("user").(*jwt.Token)
	if claims, ok := token.Claims.(*JwtCustomClaims); ok {
		return UserContext{
			Id: claims.Id,
			Name: claims.Name,
		}
	}
	return UserContext{}
}