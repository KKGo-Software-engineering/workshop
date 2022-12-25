package account

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name       string
		sqlFn      func() (*sql.DB, error)
		reqBody    string
		wantStatus int
		wantBody   string
	}{
		{"create account succesfully",
			func() (*sql.DB, error) {
				db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				if err != nil {
					return nil, err
				}
				row := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(cStmt).WithArgs(1000.0).WillReturnRows(row)
				return db, err
			},
			`{"balance": 1000.0}`,
			http.StatusCreated,
			`{"id": 1, "balance": 1000.0}`,
		},
		{"create with bad request",
			func() (*sql.DB, error) {
				return nil, nil
			},
			`{"ba}`,
			http.StatusBadRequest,
			`{"error": "bad request body"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			db, err := tc.sqlFn()
			h := New(db)

			// Assertions
			assert.NoError(t, err)
			if assert.NoError(t, h.Create(c)) {
				assert.Equal(t, tc.wantStatus, rec.Code)
				assert.JSONEq(t, tc.wantBody, rec.Body.String())
			}
		})
	}
}

func TestCreateAccount_Error(t *testing.T) {
	someErr := errors.New("some random error")
	tests := []struct {
		name    string
		sqlFn   func() (*sql.DB, error)
		reqBody string
		wantErr error
	}{
		{"create account succesfully",
			func() (*sql.DB, error) {
				db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				if err != nil {
					return nil, err
				}
				mock.ExpectQuery(cStmt).WithArgs(1000.0).WillReturnError(someErr)
				return db, err
			},
			`{"balance": 1000.0}`,
			someErr,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			db, err := tc.sqlFn()
			h := New(db)

			// Assertions
			assert.NoError(t, err)
			assert.ErrorIs(t, someErr, h.Create(c))
		})
	}
}
