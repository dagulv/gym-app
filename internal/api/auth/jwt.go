package auth

import (
	"net/http"
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/dagulv/gym-app/internal/model"
	"github.com/dagulv/gym-app/internal/db"
)



type Credentials struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	cred := Credentials{}
	user := model.User{}
	// return c.JSON(http.StatusBadRequest, c.Bind(&user))

	if err := c.Bind(&cred); err != nil {
		return c.JSON(http.StatusOK, err)
	}
	
	if err := db.DBConn.QueryRow("SELECT id, name, password FROM users WHERE name = ?", cred.Name).Scan(&user.Id, &user.Name, &user.Password); err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		return err
	}
	
	// Set custom claims
	claims := &model.JwtCustomClaims{
		user.Id,
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	setRefreshToken(c, t)

	return c.JSON(http.StatusOK, t)
}

func Signup(c echo.Context) error {
	cred := Credentials{}

	if err := c.Bind(&cred); err != nil {
		return c.JSON(http.StatusOK, err)
	}
	err := db.DBConn.QueryRow("SELECT name FROM users WHERE name = ?", cred.Name).Scan(&cred.Name)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cred.Password), bcrypt.DefaultCost); 
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Server error, unable to create your account.")
		}

		_, err = db.DBConn.Exec("INSERT INTO users(name, password, timeCreated, timeUpdated) VALUES(?, ?, ?, ?)", cred.Name, hashedPassword, time.Now(), time.Now())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Server error, unable to create your account.")
		}
	case err != nil:
		return echo.NewHTTPError(http.StatusInternalServerError, "Server error, unable to create your account.")
	}
	return nil
}