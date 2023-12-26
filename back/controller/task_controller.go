package controller

import (
	"back/db"
	"back/model"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}
	task := model.Task{}
	result := db.DB.First(&task, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Task not found")
	} else if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}
	newTaskData := model.Task{}
	if err := c.Bind(&newTaskData); err != nil {
		return err
	}
	newTaskData.UpdatedAt = time.Now()
	result := db.DB.Model(&model.Task{}).Where("id = ?", id).Updates(newTaskData)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Task not found")
	} else if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	updatedTask := model.Task{}
	if err := db.DB.First(&updatedTask, id).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}
	task := model.Task{}
	result := db.DB.First(&task, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Task not found")
	} else if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	result = db.DB.Delete(&model.Task{}, id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting task")
	}

	return c.JSON(http.StatusOK, task)
}
