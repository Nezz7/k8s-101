package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", func(c echo.Context) error {
		name := os.Getenv("NAME")
		return c.HTML(http.StatusOK, fmt.Sprintf("Hello, %s!", name))
	})

	e.GET("/version", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "v1.0.0"})
	})

	e.GET("/secret", func(c echo.Context) error {
		secret, err := os.OpenFile("/var/secrets/secret.yaml", os.O_RDONLY, 0644)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Something went wrong")
		}
		defer secret.Close()

		buf := make([]byte, 1024)
		n, err := secret.Read(buf)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Something went wrong")
		}
		return c.String(http.StatusOK, string(buf[:n]))
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
