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

	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/kkgoo-software-engineering/workshop/internal/middleware/authmw"
	"github.com/kkgoo-software-engineering/workshop/internal/middleware/zapmw"
	"github.com/kkgoo-software-engineering/workshop/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.New().All()
	e := echo.New()
	gCtx := context.Background()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		logger.Fatal("unable to configure database", zap.Error(err))
	}

	logmw := zapmw.New(logger)
	e.Use(logmw)

	authmw := authmw.Authenicate()
	e.Use(middleware.BasicAuth(authmw))

	router.RegRoute(&cfg, e, sql)

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

	ctx, cancel := context.WithTimeout(gCtx, 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal("unexpected shutdown the server", zap.Error(err))
	}
}
