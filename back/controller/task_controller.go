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

func GetTasks(c echo.Context) error {
	tasks := []model.Task{}
	db.DB.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

func GetTask(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	db.DB.Take(&task)
	return c.JSON(http.StatusOK, task)
}

// func UpdateTask(c echo.Context) error {
// 	task := model.Task{}
// 	if err := c.Bind(&task); err != nil {
// 		return err
// 	}
// 	db.DB.Save(&task)
// 	return c.JSON(http.StatusOK, task)
// }

func DeleteTask(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	db.DB.Delete(&task)
	return c.JSON(http.StatusOK, task)
}
