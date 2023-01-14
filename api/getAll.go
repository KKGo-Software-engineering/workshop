package api

import (
	"database/sql"
	"net/http"

	"github.com/kkgo-software-engineering/workshop/database"
	"github.com/labstack/echo"
)

type handler struct {
	db *sql.DB
}

func New(db *sql.DB) *handler {
	return &handler{db}
}

func (h handler) GetAllPockets(c echo.Context) error {

	_, err := database.GetAllPockets(h.db)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "hey Gopher!, I'm alive!")
}
