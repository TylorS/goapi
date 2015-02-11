package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("Route: %s, Method: %s, Processed In: %s", r.URL.Path, r.Method, t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}
