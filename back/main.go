package main

import (
	"back/controller"
	"back/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	database, _ := db.DB.DB()
	defer database.Close()

	e.GET("/tasks", controller.GetTasks)
	e.GET("/task/:id", controller.GetTask)
	e.POST("/tasks", controller.CreateTask)
	e.PUT("/task/:id", controller.UpdateTask)
	e.DELETE("/task/:id", controller.DeleteTask)
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
