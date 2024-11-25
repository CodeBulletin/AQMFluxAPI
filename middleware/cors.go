package middleware

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/config"
)

func AllowCors() Middleware {
	config.GetAPIConfig().Load()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", config.GetAPIConfig().Origins())
			w.Header().Set("Access-Control-Allow-Methods", config.GetAPIConfig().Methods())
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Refresh-Token")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			next.ServeHTTP(w, r)
		})
	}
}
