package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kkgoo-software-engineering/workshop/handler/featflag"
	"github.com/kkgoo-software-engineering/workshop/internal/config"
	"github.com/kkgoo-software-engineering/workshop/internal/middleware/zapmw"
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	hFeatFlag := featflag.New(&cfg)

	e.GET("/features", hFeatFlag.List)

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
