package zapmw

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const cZapLogger = "ZapLogger"

func New(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return zapLogMiddleware(next, logger)
	}
}

func zapLogMiddleware(next echo.HandlerFunc, logger *zap.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Set(cZapLogger, logger)
		next(c)
		return nil
	}
}

func Logger(c echo.Context) *zap.Logger {
	logger, ok := c.Get(cZapLogger).(*zap.Logger)
	if !ok {
		logger = zap.NewNop()
	}
	return logger
}
