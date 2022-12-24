package featflag

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/kkgoo-software-engineering/workshop/internal/config"
)

func List(cfg *config.Config) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, cfg.FeatureFlag)
	}
}
