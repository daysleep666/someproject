package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/request", func(c echo.Context) error {
		return c.Redirect(302, "https://www.baidu.com")
	})

	err := e.Start(":2234")
	if err != nil {
		panic(err)
	}
}
