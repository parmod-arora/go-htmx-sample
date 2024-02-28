package home

import (
	"net/http"
	"parmod/go-htmx-sample/infrastructure/loglib"
	"parmod/go-htmx-sample/internal/application"
	"parmod/go-htmx-sample/views"
	"strconv"
)

func RootHandler(app *application.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := loglib.GetLogger(r.Context())
		index := views.Index("Title here")

		w.Header().Set("Content-Type", "text/html charset=utf-8")
		err := views.Layout(index, "ToDo Page").Render(r.Context(), w)
		if err != nil {
			logger.Errorf("Error while rendering layout %s", err.Error())
		}
	}
}

func GetTodos(app *application.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := loglib.GetLogger(r.Context())
		todoItemView := views.TodoItem("Todo Item")
		w.Header().Set("Content-Type", "text/html charset=utf-8")
		if !isHtmxRequest(r) {
			err := views.Layout(todoItemView, "ToDo Item Page").Render(r.Context(), w)
			if err != nil {
				logger.Errorf("Error while rendering layout %s", err.Error())
			}
			return
		}
		err := todoItemView.Render(r.Context(), w)
		if err != nil {
			logger.Errorf("Error while rendering layout %s", err.Error())
		}
	}
}

func isHtmxRequest(r *http.Request) bool {
	hx_request, err := strconv.ParseBool(r.Header.Get("Hx-Request"))
	if err != nil {
		return false
	}
	return hx_request
}
