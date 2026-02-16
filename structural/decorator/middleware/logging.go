package middleware

import (
	"log"
	"time"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		recorder := NewStatusRecorder(w, http.StatusOK)

		next.ServeHTTP(recorder, r)

		// Log the result
		log.Printf("[METHOD] %s | [PATH] %s | [STATUS] %d | [DURATION] %v",
			r.Method, r.URL.Path, recorder.StatusCode, time.Since(start))
	})
}