//go:build integration

package account

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountIT(t *testing.T) {
	e := echo.New()

	cfg := config.New().All()
	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		t.Error(err)
	}

	hAccount := New(sql)

	e.POST("/accounts", hAccount.Create)

	reqBody := `{"balance": 999.99}`
	req := httptest.NewRequest(http.MethodPost, "/accounts", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	expected := `{"id": 1, "balance": 999.99}`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expected, rec.Body.String())
}
