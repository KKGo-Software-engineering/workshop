package router

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// this one probably not necessarily to test, it's really straight forward
func TestRegisterRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	RegRoute(&config.Config{}, e, &sql.DB{})
	rts := e.Routes()
	assert.Equal(t, len(rts), 4)
}
