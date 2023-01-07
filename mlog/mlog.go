package mlog

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const key = "logger"

func L(c echo.Context) *zap.Logger {
	switch logger := c.Get(key).(type) {
	case *zap.Logger:
		return logger
	default:
		return zap.NewNop()
	}
}
