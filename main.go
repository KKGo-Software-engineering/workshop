package main

import (
	"fmt"
	"net/http"

	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	cfg := config.New().All()
	addr := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)
	e.Logger.Fatal(e.Start(addr))
}
