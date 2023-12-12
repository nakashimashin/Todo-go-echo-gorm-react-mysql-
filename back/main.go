package main

import (
	"back/controller"
	"back/db"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database, _ := db.DB.DB()
	defer database.Close()

	e.GET("/tasks", controller.GetTasks)
	e.GET("/tasks/:id", controller.GetTask)
	e.POST("/tasks", controller.CreateTask)
	// e.PUT("/tasks/:id", controller.UpdateTask)
	e.DELETE("/tasks/:id", controller.DeleteTask)
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
