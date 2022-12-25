package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/kkgoo-software-engineering/workshop/internal/middleware/zapmw"
	"github.com/kkgoo-software-engineering/workshop/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New().All()
	e := echo.New()
	gCtx := context.Background()

	logger, _ := zap.NewProduction()
	logmw := zapmw.New(logger)
	e.Use(logmw)

	router.RegRoute(&cfg, e)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)

	go func() {
		err := e.Start(addr)
		if err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
		e.Logger.Print("gracefully shutdown the server")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(gCtx, 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
