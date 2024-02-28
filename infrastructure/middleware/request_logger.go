package middleware

import (
	"net/http"
	"parmod/go-htmx-sample/infrastructure/loglib"
	"strings"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		start := time.Now()
		ctx := r.Context()

		// Parse request information
		requestURIparts := append(strings.SplitN(r.RequestURI, "?", 2), "") // `append` so we'd always have an array of 2 strings at least
		rctx := NewRequestContext(r)
		ctx = SetRequestContext(ctx, rctx)

		loglib.DefaultConciseLogger()

		// Instantiate verbose logger
		logger := loglib.DefaultConciseLogger().
			WithField("request", rctx.RequestID).
			WithField("route", r.Method+" "+requestURIparts[0]).
			WithField("query", requestURIparts[1]).
			WithField("ip", r.RemoteAddr).
			WithField("referer", r.Referer()).
			WithField("agent", r.UserAgent())

		// Installation ID
		if installationID := r.Header.Get("X-Installation-ID"); installationID != "" {
			logger = logger.WithField("installation", installationID)
		}

		// From user agent
		logger = logger.
			WithField("useragent", r.UserAgent())

		// Set logger into context
		ctx = loglib.SetLogger(ctx, logger)

		logger.
			Infof("START")

		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)

		logger.
			WithField("duration", time.Since(start)).
			Infof("END")
	})
}
