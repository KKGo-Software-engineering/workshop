package router

import (
	"database/sql"
	"net/http"

	"github.com/kkgo-software-engineering/workshop/handler/account"
	"github.com/kkgo-software-engineering/workshop/handler/featflag"
	"github.com/kkgo-software-engineering/workshop/handler/healthchk"
	"github.com/kkgo-software-engineering/workshop/internal/config"
	"github.com/labstack/echo/v4"
)

func RegRoute(cfg config.Config, e *echo.Echo, db *sql.DB) {
	hHealthChk := healthchk.New(db)
	e.GET("/healthz", hHealthChk.Check) // TODO: did help need auth?

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	hAccount := account.New(cfg.FeatureFlag, db)
	e.POST("/accounts", hAccount.Create)

	hFeatFlag := featflag.New(cfg)
	e.GET("/features", hFeatFlag.List)
}
