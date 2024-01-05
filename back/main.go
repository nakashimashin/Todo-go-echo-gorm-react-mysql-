package main

import (
	"back/controller"
	"back/db"
	"net/http"
	"os"

	"log"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in .env file")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	database, err := db.DB.DB()
	if err != nil {
		log.Fatalf("Error getting datsbase connection: %v", err)
	}
	defer database.Close()

	e.POST("/signup", controller.SignUp)
	e.POST("/login", controller.Login)
	e.POST("/logout", controller.Logout)

	jwtMiddleware := echojwt.JWT([]byte(jwtSecret))

	e.GET("/tasks", controller.GetTasks, jwtMiddleware)
	e.GET("/task/:id", controller.GetTask, jwtMiddleware)
	e.POST("/task", controller.CreateTask, jwtMiddleware)
	e.PUT("/task/:id", controller.UpdateTask, jwtMiddleware)
	e.DELETE("/task/:id", controller.DeleteTask, jwtMiddleware)
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
