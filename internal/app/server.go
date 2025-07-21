package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *Application) Serve(routes http.Handler) error {
	srv := &http.Server{
		Addr:         app.Config.Port,
		Handler:      routes,
		IdleTimeout:  app.Config.IdleTimeout,
		ReadTimeout:  app.Config.ReadTimeout,
		WriteTimeout: app.Config.WriteTImeout,
		ErrorLog:     log.New(os.Stderr, "HTTP_ERROR", log.LstdFlags),
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// block until signal is received
		s := <-quit
		app.Logger.Printf("Shutting down server, signal: %s\n", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		app.Logger.Printf("Starting graceful shutdown (30s timeout)")
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
			return
		}

		app.Logger.Printf("Waiting for background tasks to complete...\n")
		app.WG.Wait()
		app.Logger.Printf("All background tasks completed\n")

		shutdownError <- nil
	}()

	app.Logger.Printf("Starting server on %s (env: %s)\n", srv.Addr, app.Config.AppEnv)

	err := srv.ListenAndServe()

	// check if this is the expected "server closed" error
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// wait for shutdown to complete and get any shutdowns errors
	err = <-shutdownError
	if err != nil {
		return err
	}

	app.Logger.Printf("Server stopped gracefully")
	return nil
}
