package middleware

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Log request
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Call the actual handler
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
