package middleware

import "net/http"

func DefaultHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
