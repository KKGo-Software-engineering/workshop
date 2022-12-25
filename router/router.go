package router

import (
	"database/sql"
	"net/http"

	"github.com/kkgoo-software-engineering/workshop/handler/featflag"
	"github.com/kkgoo-software-engineering/workshop/handler/healthchk"
	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/labstack/echo/v4"
)

func RegRoute(cfg *config.Config, e *echo.Echo, db *sql.DB) {
	hHealthChk := healthchk.New(db)
	e.GET("/healthz", hHealthChk.Check)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	hFeatFlag := featflag.New(cfg)
	e.GET("/features", hFeatFlag.List)
}
