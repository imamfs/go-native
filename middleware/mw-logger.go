package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip := r.RemoteAddr
		next.ServeHTTP(w, r)
		fmt.Printf("method=%s path=%s address= %s duration=%s\n", r.Method, r.URL.Path, ip, time.Since(start))
	})
}
