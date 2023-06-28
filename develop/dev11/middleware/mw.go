package middleware

import (
	"log"
	"net/http"
)

func RequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("%s %s", r.Method, r.RequestURI)
	})
}
