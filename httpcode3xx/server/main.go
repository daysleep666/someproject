package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/request", func(c echo.Context) error {
		return c.Redirect(302, "test")
	})

	e.GET("/test", func(c echo.Context) error {
		return c.Redirect(302, "test")
	})

	err := e.Start(":2234")
	if err != nil {
		panic(err)
	}
}
