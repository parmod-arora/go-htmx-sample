package router

import (
	"net/http"
	"parmod/go-htmx-sample/infrastructure/middleware"
	"parmod/go-htmx-sample/internal/application"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handler(app *application.App) http.Handler {

	router := mux.NewRouter()
	router.StrictSlash(true)
	router.Use(middleware.Recover)

	// Register prometheus
	middleware.RegisterPrometheus()
	router.Use(middleware.PrometheusMiddleware)
	router.Path("/prometheus").Handler(promhttp.Handler())

	// This will serve files under http://localhost:8000/static/<filename>
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", CacheControlWrapper(http.FileServer(http.Dir("./static")))))
	router.Handle("/", http.RedirectHandler("/todos/", http.StatusSeeOther))

	// app pages router
	todosRouter(app, router)

	// api router
	apiRoutes(app, router)
	return router
}

func CacheControlWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=2592000") // 30 days
		h.ServeHTTP(w, r)
	})
}
