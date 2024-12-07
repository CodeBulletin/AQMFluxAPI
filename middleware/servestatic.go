package middleware

import (
	"net/http"
)

func ServeStatic (fs http.FileSystem) func(http.Handler) http.Handler {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				next.ServeHTTP(w, r)
				return
			}

			file, err := fs.Open(r.URL.Path)
			if err != nil {
				file, err = fs.Open("index.html")
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
			}

			defer file.Close()

			stat, err := file.Stat()
			if err != nil {
				file, err = fs.Open("index.html")
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				defer file.Close()

				stat, err = file.Stat()
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
			}

			if stat.IsDir() {
				file, err = fs.Open("index.html")
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				defer file.Close()

				stat, err = file.Stat()
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.ServeContent(w, r, stat.Name(), stat.ModTime(), file)
		})
	}
}