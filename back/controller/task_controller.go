package controller

import (
	"back/db"
	"back/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTask(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	db.DB.Create(&task)
	return c.JSON(http.StatusCreated, task)
}
