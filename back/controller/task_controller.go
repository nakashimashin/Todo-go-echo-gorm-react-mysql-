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
	// URLパラメータからIDを取得
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// IDが無効な場合はエラーを返す
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	// 指定されたIDのタスクを取得
	task := model.Task{}
	result := db.DB.First(&task, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// タスクが見つからない場合はエラーを返す
		return echo.NewHTTPError(http.StatusNotFound, "Task not found")
	} else if result.Error != nil {
		// その他のデータベースエラーがある場合はエラーを返す
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	// タスクをJSON形式で返す
	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	// URLパラメータからIDを取得
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// IDが無効な場合はエラーを返す
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	// タスクの新しい内容を取得
	newTaskData := model.Task{}
	if err := c.Bind(&newTaskData); err != nil {
		return err
	}

	// UpdatedAtを現在時刻に設定
	newTaskData.UpdatedAt = time.Now()

	// 指定されたIDのタスクを更新
	result := db.DB.Model(&model.Task{}).Where("id = ?", id).Updates(newTaskData)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// タスクが見つからない場合はエラーを返す
		return echo.NewHTTPError(http.StatusNotFound, "Task not found")
	} else if result.Error != nil {
		// その他のデータベースエラーがある場合はエラーを返す
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	// 更新されたタスクを取得して返す
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
