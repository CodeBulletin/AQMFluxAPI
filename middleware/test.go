package middleware

import (
	"net/http"
	"time"
)

func Slow(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		// write an error message to the response
		// utils.WriteError(w, http.StatusRequestTimeout, fmt.Errorf("request timed out"))
		next.ServeHTTP(w, r)
	})
}
