//go:build unit

package mlog

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLogMiddleware(t *testing.T) {
	e := echo.New()
	// error can be ignore IF it's a test.
	logger, _ := zap.NewProduction()
	mw := New(logger)
	e.Use(mw)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	ctx := e.AcquireContext()
	assert.IsType(t, &zap.Logger{}, Logger(ctx))
}

func TestUnsetLogMiddleware(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	ctx := e.AcquireContext()
	assert.IsType(t, &zap.Logger{}, Logger(ctx))
}
