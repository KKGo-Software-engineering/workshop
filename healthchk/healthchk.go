package healthchk

import (
	"database/sql"
	"net/http"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
)

type handler struct {
	db *sql.DB
}

func New(db *sql.DB) *handler {
	return &handler{db}
}

func (h handler) Check(c echo.Context) error {
	logger := mlog.L(c)
	logger.Info("health check")
	err := h.db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "hey Gopher!, I'm alive!")
}
