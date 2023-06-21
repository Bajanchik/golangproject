package mw

import (
	"log"

	"github.com/labstack/echo/v4"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		value := c.Request().Header.Get("User-Role")

		if value == "admin" {
			log.Println("red button user detected")
		}

		err := next(c)
		if err != nil {
			return err
		}

		return nil
	}
}
