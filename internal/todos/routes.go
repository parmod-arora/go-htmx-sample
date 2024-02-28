package todos

import (
	"encoding/json"
	"net/http"
	"parmod/go-htmx-sample/internal/application"

	"github.com/gorilla/mux"
)

type RouteHandler func(http.ResponseWriter, *http.Request)

func ApiRoutes(app *application.App, rootRouter *mux.Router) {
	// api routes
	r := rootRouter.PathPrefix("/todos").Subrouter()
	r.HandleFunc("/", ReadTodo(app)).Methods("GET")
}

func ReadTodo(app *application.App) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		todos := []Todo{
			{
				Task:      "task1",
				Completed: false,
			},
		}

		bytes, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
