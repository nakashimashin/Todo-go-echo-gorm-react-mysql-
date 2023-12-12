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
	// URLパラメータからIDを取得
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// IDが無効な場合はエラーを返す
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	// 削除するタスクを取得
	task := model.Task{}
	result := db.DB.First(&task, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// タスクが見つからない場合はエラーを返す
		return echo.NewHTTPError(http.StatusNotFound, "Task not found")
	} else if result.Error != nil {
		// その他のデータベースエラーがある場合はエラーを返す
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	// タスクを削除
	result = db.DB.Delete(&model.Task{}, id)
	if result.Error != nil {
		// 削除中にエラーが発生した場合は、エラーを返す
		return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting task")
	}

	// 削除されたタスクの内容をJSON形式で返す
	return c.JSON(http.StatusOK, task)
}
