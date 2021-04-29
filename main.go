package main

import (
	"github.com/labstack/echo/v4"
	"github.com/naokivandit/handson-echo-go/routes"
)

func main() {
	e := echo.New()
	routes.Routes(e)
	e.Start(":8000")
}
