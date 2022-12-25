package healthchk

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {

	tests := []struct {
		name  string
		sqlFn func() (*sql.DB, error)
		want  int
	}{
		{"all healthy",
			func() (*sql.DB, error) {
				db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
				if err != nil {
					return nil, err
				}
				mock.ExpectPing().WillReturnError(nil)
				return db, err
			},
			http.StatusOK},
		{"ping return error",
			func() (*sql.DB, error) {
				db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
				if err != nil {
					return nil, err
				}
				mock.ExpectPing().WillReturnError(errors.New("some unexpected error"))
				return db, err
			},
			http.StatusInternalServerError},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			db, err := tc.sqlFn()
			h := New(db)

			// Assertions
			assert.NoError(t, err)
			if assert.NoError(t, h.Check(c)) {
				assert.Equal(t, tc.want, rec.Code)
			}
		})
	}
}
