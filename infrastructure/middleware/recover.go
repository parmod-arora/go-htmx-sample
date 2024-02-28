package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"parmod/go-htmx-sample/infrastructure/loglib"
	"runtime"
)

func Recover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger := loglib.GetLogger(r.Context())
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				fmt.Printf("recovering from err %v\n %s", err, buf)
				logger.Errorf("recovering from err %s", err)

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
