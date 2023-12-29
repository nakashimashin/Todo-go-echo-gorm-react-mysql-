package controller

import (
	"back/db"
	"back/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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

	responseUser := model.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return c.JSON(http.StatusCreated, responseUser)
}

func Login(c echo.Context) error {
	creds := model.LoginCredentials{}
	if err := c.Bind(&creds); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	user := model.User{}
	if result := db.DB.Where("email = ?", creds.Email).First(&user); result.Error != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("your_jwt_secret"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
