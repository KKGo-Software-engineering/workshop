package api

import (
	"database/sql"
	"net/http"

	"github.com/kkgo-software-engineering/workshop/database"
	"github.com/labstack/echo"
)

func GetAllPockets(c echo.Context, db *sql.DB) error {

	exps, err := database.GetAllPockets(db)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, exps)
}
