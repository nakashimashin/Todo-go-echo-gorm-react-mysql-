package controller

import (
	"back/db"
	"back/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Password hashing failed")
	}
	user.Password = string(hashedPassword)

	if result := db.DB.Create(&user); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	// passwordを空にして返す
	responseUser := model.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return c.JSON(http.StatusCreated, responseUser)
}
