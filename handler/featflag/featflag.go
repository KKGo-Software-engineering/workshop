package featflag

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/kkgo-software-engineering/workshop/internal/config"
	"github.com/kkgo-software-engineering/workshop/internal/middleware/zapmw"
)

type handler struct {
	cfg *config.Config
}

func New(cfg *config.Config) *handler {
	return &handler{cfg}
}

func (h handler) List(c echo.Context) error {
	logger := zapmw.Logger(c)
	defer logger.Sync()

	logger.Info("called api", zap.String("test-key", "test-value"))
	return c.JSON(http.StatusOK, h.cfg.FeatureFlag)
}
