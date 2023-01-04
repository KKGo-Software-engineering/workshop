//go:build unit

package router

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kkgo-software-engineering/workshop/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// this one probably not necessarily to test, it's really straight forward
func TestRegisterRoute(t *testing.T) {
	e := RegRoute(config.Config{}, zap.NewNop(), &sql.DB{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	rts := e.Routes()

	assert.Equal(t, len(rts), 4)
}
