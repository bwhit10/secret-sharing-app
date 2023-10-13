package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Renderer = &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	// Routes
	app.Static("/assets", "./assets")
	app.GET("/", index)
	app.POST("/save", save)

	app.Logger.Fatal(app.Start(":3000"))
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func save(c echo.Context) error {
	// TODO: save c.FormValue to display later

	return c.String(http.StatusOK, "TODO")
}
