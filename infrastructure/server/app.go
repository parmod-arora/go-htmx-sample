package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"parmod/go-htmx-sample/infrastructure/loglib"
	"syscall"
	"time"
)

// App is a server that handles HTTP requests
type App struct {
	server *http.Server
	logger loglib.Logger
}

// New creates an App server instance
func New(addr string, handler http.Handler) *App {
	return &App{
		server: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		logger: loglib.DefaultConciseLogger(),
	}
}

// Start starts the application server
func (a *App) Start() {
	// Start server asynchronously
	go func() {
		a.logger.Infof("Server started at port %s", a.server.Addr)
		if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
			a.logger.Fatalf("ListenAndServe: %s", err)
		}
	}()

	// Handle graceful shutdown
	// Channel to listen for an interrupt or terminate signal from the OS.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	// Block waiting for a receive on signal from OS
	s := <-osSignals
	switch s {
	case syscall.SIGTERM:
		d := 10 * time.Second
		a.logger.Infof("SIGTERM received. Sleeping for %s as buffer before stopping server", d)
		// Delay 10 seconds as buffer
		// i.e. 'time kubectl delete endpoints {name}' for cluster is about 5 seconds
		time.Sleep(d)
	}

	// Shutdown gracefully
	a.Stop()
}

// Stop stops the application server
func (a *App) Stop() {
	// Create a context to attempt a graceful 5 second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Attempt the graceful shutdown by closing the listener and
	// completing all inflight requesta.
	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Errorf("Could not stop server gracefully: %v", err)
		a.logger.Infof("Initiating hard shutdown")
		if err := a.server.Close(); err != nil {
			a.logger.Errorf("Could not stop http server: %v", err)
		}
	}
}
