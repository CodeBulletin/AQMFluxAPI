package middleware

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/codebulletin/AQMFluxAPI/logger"
)

type responseWriter struct {
	http.ResponseWriter;
	status 				int;
	wroteHeader 		bool;
}

func (rw *responseWriter) Status() int {
	return rw.status
}
  
func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func Logger(logger logger.Logger) func(http.Handler) http.Handler {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fn := func(w http.ResponseWriter, r *http.Request) {
				defer func() {
					if err := recover(); err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						// log.Printf(
						// 	"err: %v\ntrace: %v",
						// 	err,
						// 	debug.Stack(),
						// )
						logger.Error("err: %v\ntrace: %v", err, debug.Stack())
					}
				}()
				start := time.Now()
				ww := wrapResponseWriter(w)
				next.ServeHTTP(ww, r)
				// log.Printf("%d %s %s %s %s", ww.status, r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
				logger.Info("%d %s %s %s %s", ww.status, r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
			}

			fn(w, r)
		})
	}
}