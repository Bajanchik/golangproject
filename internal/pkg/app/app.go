package app

import (
	"fmt"
	"golangproject/internal/app/endpoint"
	"golangproject/internal/app/endpoint/service"
	"golangproject/internal/app/mw"
	"log"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server running")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
