package controller

import (
	"back/db"
	"back/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getUserIDFromToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(uint)
	return userID
}

func CreateTask(c echo.Context) error {
	userID := getUserIDFromToken(c)

	task := model.Task{UserID: userID}
	if err := c.Bind(&task); err != nil {
		return err
	}
	db.DB.Create(&task)
	return c.JSON(http.StatusCreated, task)
}

func GetTasks(c echo.Context) error {
	userID := getUserIDFromToken(c)

	tasks := []model.Task{}
	db.DB.Where("user_id = ?", userID).Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

func GetTask(c echo.Context) error {
	userID := getUserIDFromToken(c)
	id, _ := strconv.Atoi(c.Param("id"))
	var task model.Task
	result := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&task)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Task not found")
	} else if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	userID := getUserIDFromToken(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	var existingTask model.Task
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&existingTask).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Task not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	updateData := model.Task{}
	if err := c.Bind(&updateData); err != nil {
		return err
	}
	updateData.ID = existingTask.ID
	updateData.UserID = existingTask.UserID

	if err := db.DB.Model(&existingTask).Updates(updateData).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating task")
	}

	return c.JSON(http.StatusOK, updateData)
}

func DeleteTask(c echo.Context) error {
	userID := getUserIDFromToken(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	var task model.Task
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Task not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := db.DB.Delete(&model.Task{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting task")
	}

	return c.JSON(http.StatusOK, task)
}
