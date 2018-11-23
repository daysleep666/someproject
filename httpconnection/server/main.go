package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-store")
			return next(c)
		}
	})
	e.Static("/hello", "httpcachecontrol/view")
	e.File("/hello", "httpcachecontrol/view/*.html")
	err := e.Start(":2234")
	if err != nil {
		panic(err)
	}
}
