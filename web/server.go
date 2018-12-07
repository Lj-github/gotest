package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/aaa", func(c echo.Context) error {
		return c.String(http.StatusOK, "aaaaaaaaaa!")
	})
	e.Logger.Fatal(e.Start(":1323"))
	print("localhost:")
}
