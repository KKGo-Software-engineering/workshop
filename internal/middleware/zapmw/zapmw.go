package zapmw

import (
	"github.com/google/uuid"
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
		ehLogger := ehLogger(c, logger)
		c.Set(cZapLogger, ehLogger)
		next(c)
		return nil
	}
}

func ehLogger(c echo.Context, logger *zap.Logger) *zap.Logger {
	xParent := c.Request().Header.Get("X-Parent-ID")
	if xParent == "" {
		xParent = uuid.NewString()
	}
	xSpan := uuid.NewString()
	return logger.With(zap.String("parent-id", xParent),
		zap.String("span-id", xSpan))
}

func Logger(c echo.Context) *zap.Logger {
	logger, ok := c.Get(cZapLogger).(*zap.Logger)
	if !ok {
		logger = zap.NewNop()
	}
	return logger
}
