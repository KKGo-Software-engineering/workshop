package featflag

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/kkgoo-software-engineering/workshop/internal/config"
)

type handler struct {
	cfg *config.Config
}

func New(cfg *config.Config) *handler {
	return &handler{cfg}
}

func (h handler) List(c echo.Context) error {
	return c.JSON(http.StatusOK, h.cfg.FeatureFlag)
}
