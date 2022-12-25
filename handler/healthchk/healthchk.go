package healthchk

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	db *sql.DB
}

func New(db *sql.DB) *handler {
	return &handler{db}
}

func (h handler) Check(c echo.Context) error {
	err := h.db.Ping()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
