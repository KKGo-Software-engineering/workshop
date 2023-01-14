package api

import (
	"database/sql"
	"net/http"

	"github.com/kkgo-software-engineering/workshop/database"
	"github.com/labstack/echo/v4"
)

type handler struct {
	db *sql.DB
}

func New(db *sql.DB) *handler {
	return &handler{db}
}

func (h handler) GetAllPockets(c echo.Context) error {

	id := c.Param("id")
	pockets, err := database.GetAllPockets(h.db, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, pockets)
}
