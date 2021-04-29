package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naokivandit/handson-echo-go/domain"
)

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func Get(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
