package application

import (
	"github.com/go-chi/chi/v5"
	slogchi "github.com/samber/slog-chi"
	"log/slog"
	"net/http"
	"notificationService/internal/receivers"
	"notificationService/internal/routes"
)

type App struct {
	router *chi.Mux
}

// New creates a new app instance
// For logging slog.Default() is used
func New() *App {
	router := chi.NewRouter()
	router.Use(slogchi.New(slog.Default()))

	return &App{
		router: router,
	}
}

// Run is the main entry point for running the notification service
// It registers all Receivers listed under folder 'receivers'
// In development mode the application is listening on localhost port 3000
func (instance *App) Run() error {
	err := instance.UseAndRun([]func(*chi.Mux){func(router *chi.Mux) {
		routes.UseNotificationsRoute(router, receivers.NewReceivers())
	}})
	return err
}

// UseAndRun runs the notification service with a list of custom set of middleware functions
// Mostly used for integration tests where a custom setup is needed
func (instance *App) UseAndRun(middlewareFunctions []func(router *chi.Mux)) error {
	for _, middlewareFunction := range middlewareFunctions {
		middlewareFunction(instance.router)
	}

	err := chi.Walk(instance.router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		slog.Info("Middleware", "Method", method, "Route", route)
		return nil
	})
	if err != nil {
		slog.Warn("Could not log middleware routes", "warning", err)
	}

	// TODO move port to config
	err = http.ListenAndServe(":3000", instance.router)
	if err != nil {
		slog.Error("Failed to server requests", "error", err)
		return err
	}

	return nil
}
