package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Topic struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/topics", getTopics)

	e.Logger.Fatal(e.Start(":8000"))
}

func getTopics(c echo.Context) error {
	return c.JSON(http.StatusOK, []Topic{
		{ID: 1, Title: "Goの使い方"},
		{ID: 2, Title: "Echoの使い方"},
	})
}
