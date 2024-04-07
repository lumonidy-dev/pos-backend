package middleware

import (
	"log"
	"net/http"
)

// statusWriter es una estructura que envuelve http.ResponseWriter para rastrear el código de estado de la respuesta
type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader intercepta el código de estado de la respuesta
func (w *statusWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// LogStatusMiddleware registra la dirección IP del cliente y el estado de la petición en los logs
func LogStatusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := statusWriter{ResponseWriter: w}
		next.ServeHTTP(&sw, r)
		log.Printf("[%s] Solicitud recibida desde %s: %s %s", http.StatusText(sw.statusCode), r.RemoteAddr, r.Method, r.URL.Path)
	})
}
