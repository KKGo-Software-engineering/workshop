package router

import (
	"net/http"

	"github.com/kkgoo-software-engineering/workshop/handler/featflag"
	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/labstack/echo/v4"
)

func RegRoute(cfg *config.Config, e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	hFeatFlag := featflag.New(cfg)
	e.GET("/features", hFeatFlag.List)
}
