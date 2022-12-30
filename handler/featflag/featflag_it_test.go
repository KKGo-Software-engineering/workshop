//go:build integration

package featflag

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kkgo-software-engineering/workshop/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetFeatFlagIT(t *testing.T) {
	e := echo.New()

	hFeatFlag := New(&config.Config{
		FeatureFlag: config.FeatureFlag{IsLimitMaxBalanceOnCreate: true}})

	e.GET("/features", hFeatFlag.List)

	req := httptest.NewRequest(http.MethodGet, "/features", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	expected := `{"isLimitMaxBalanceOnCreate": true}`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expected, rec.Body.String())
}
