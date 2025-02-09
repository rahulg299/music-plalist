// middleware/middleware.go
package middleware

import (
    "log"
    "net/http"
    "strings"
    "time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf(
            "%s %s %s",
            r.Method,
            r.RequestURI,
            time.Since(start),
        )
    })
}

func JSONContentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if the path starts with /api using strings.HasPrefix instead of slice
        if strings.HasPrefix(r.URL.Path, "/api") {
            w.Header().Add("Content-Type", "application/json")
        }
        next.ServeHTTP(w, r)
    })
}

