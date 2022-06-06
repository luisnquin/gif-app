package core

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
)

func GracefulShutdown(app *echo.Echo) (startup func(string), wait func(), shutdown func()) {
	startup = func(port string) {
		if err := app.Start(port); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal(err)
		}
	}

	wait = func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt, os.Kill)
		<-shutdown
	}

	shutdown = func() {
		if err := app.Shutdown(context.Background()); err != nil {
			app.Logger.Fatal(err)
		}
	}

	return startup, wait, shutdown
}
