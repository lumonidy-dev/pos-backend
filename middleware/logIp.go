package middleware

import (
	"log"
	"net/http"
	"pos-backend/util"
)

// LogIPMiddleware registra la dirección IP del cliente en los logs
func LogIPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := util.GetClientIP(r)
		log.Printf("Solicitud recibida desde %s: %s %s", clientIP, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
