package controller

import (
	"back/db"
	"back/model"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getUserIDFromToken(c echo.Context) (uint, error) {
	log.Println(c)
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return 0, errors.New("user not found in context")
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("error retrieving user claims from token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user ID not found in token")
	}

	return uint(userID), nil
}

func CreateTask(c echo.Context) error {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	task := model.Task{UserID: userID}
	if err := c.Bind(&task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result := db.DB.Create(&task); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusCreated, task)
}

func GetTasks(c echo.Context) error {
	userID, err := getUserIDFromToken(c)
	log.Println(err)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	var tasks []model.Task
	if result := db.DB.Where("user_id = ?", userID).Find(&tasks); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTask(c echo.Context) error {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	var task model.Task
	if result := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&task); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Task not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	var existingTask model.Task
	if result := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&existingTask); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Task not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	updateData := model.Task{}
	if err := c.Bind(&updateData); err != nil {
		return err
	}

	if err := db.DB.Model(&existingTask).Updates(updateData).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating task")
	}

	return c.JSON(http.StatusOK, updateData)
}

func DeleteTask(c echo.Context) error {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	var task model.Task
	if result := db.DB.Where("id = ? AND user_id = ?", id, userID).First(&task); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Task not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := db.DB.Delete(&model.Task{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting task")
	}

	return c.JSON(http.StatusOK, task)
}
