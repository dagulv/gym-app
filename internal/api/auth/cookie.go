package auth

import (
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

func setRefreshToken(c echo.Context, t string) {
	cookie := new(http.Cookie)
	cookie.Name = "refreshToken"
	cookie.Value = t
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}