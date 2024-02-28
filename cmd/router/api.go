package router

import (
	"encoding/json"
	"net/http"
	"parmod/go-htmx-sample/infrastructure/middleware"
	"parmod/go-htmx-sample/internal/application"
	"parmod/go-htmx-sample/internal/todos"

	"github.com/gorilla/mux"
)

func apiRoutes(app *application.App, router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.RequestLogger)
	todos.ApiRoutes(app, apiRouter)

	// Move to different file
	apiRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	apiRouter.Use(mux.CORSMethodMiddleware(apiRouter))
}
