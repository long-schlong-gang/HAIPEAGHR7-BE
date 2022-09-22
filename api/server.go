/*
 *	Handles starting and graceful shutdown of the echo server
 */

package api

import (
	"HAIPEAGHR7/util"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kpango/glg"
	"github.com/labstack/echo/v4"
)

func StartServer(e *echo.Echo) {
	if util.GlobalEnvs.Port == 0 {
		glg.Error("No port defined; proceeding with random port from Echo.")
	}

	go func() {
		err := e.Start(fmt.Sprintf(":%d", util.GlobalEnvs.Port))
		if err != nil && err != http.ErrServerClosed {
			glg.Errorf("Server stopped with error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		glg.Errorf("Server stopped with error: %v", err)
	} else {
		glg.Info("Server stopped successfully")
	}
}
