package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/callback", CallBackHandler)
	err := e.Start("10.0.0.66:1234")
	if err != nil {
		panic(err)
	}
}

func CallBackHandler(c echo.Context) error {
	fmt.Println("i'm call back")
	return c.JSON(http.StatusOK, "")
}
