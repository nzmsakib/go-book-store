package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func Demo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("Demo middlewares")

		defer func() {
			fmt.Println("Demo middlewares end")
		}()

		c.Set("demoKey", "demoString")

		return next(c)
	}
}
