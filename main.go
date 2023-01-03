package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kkgo-software-engineering/workshop/internal/config"
	"github.com/kkgo-software-engineering/workshop/internal/middleware/auth"
	"github.com/kkgo-software-engineering/workshop/internal/middleware/mlog"
	"github.com/kkgo-software-engineering/workshop/internal/router"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func server(cfg config.Config, logger *zap.Logger, sql *sql.DB) *echo.Echo {
	e := echo.New()
	e.Use(mlog.Middleware(logger))
	e.Use(middleware.BasicAuth(auth.Authenicate()))
	router.RegRoute(cfg, e, sql)
	return e
}

func main() {
	cfg := config.New().All()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		logger.Fatal("unable to configure database", zap.Error(err))
	}

	e := server(cfg, logger, sql)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)

	go func() {
		err := e.Start(addr)
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("unexpected shutdown the server", zap.Error(err))
		}
		logger.Info("gracefully shutdown the server")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gCtx := context.Background()
	ctx, cancel := context.WithTimeout(gCtx, 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal("unexpected shutdown the server", zap.Error(err))
	}
}
