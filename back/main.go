package main

import (
	"back/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func connect(c echo.Context) error {
	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
	} else {
		return c.String(http.StatusOK, "DB接続しました")
	}
}

func main() {
	e := echo.New()
	e.GET("/", connect)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
