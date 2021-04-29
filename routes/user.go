package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/naokivandit/handson-echo-go/handler"
)

func Routes(e *echo.Echo) {
	e.GET("/", handler.HelloWorld)
	e.GET("/users/:user_id", handler.Get)
}
