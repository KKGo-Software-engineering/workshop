package featflag

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetFeatFlagIT(t *testing.T) {
	e := echo.New()

	hFeatFlag := New(&config.Config{
		FeatureFlag: config.FeatureFlag{IsLimitMaxSpend: true}})

	e.GET("/features", hFeatFlag.List)

	req := httptest.NewRequest(http.MethodGet, "/features", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	expected := `{"IsLimitMaxSpend": true}`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expected, rec.Body.String())
}
