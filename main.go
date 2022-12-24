package main

import (
	"fmt"
	"net/http"

	"github.com/kkgoo-software-engineering/workshop/handler/featflag"
	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.New().All()
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/features", featflag.List(&cfg))

	addr := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)
	e.Logger.Fatal(e.Start(addr))
}
