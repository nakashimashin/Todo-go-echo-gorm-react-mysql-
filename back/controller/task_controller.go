package controller

import (
	"back/db"
	"back/model"
	"net/http"
	"time"

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

func UpdateTask(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}

	if err := db.DB.Model(&task).Updates(model.Task{Title: task.Title, UpdatedAt: time.Now()}).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func DeleteTask(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	db.DB.Delete(&task)
	return c.JSON(http.StatusOK, task)
}
