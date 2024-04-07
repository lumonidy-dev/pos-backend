package middleware

import (
	"net/http"
	"strings"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Configurar las opciones CORS
		allowedOrigins := []string{"https://gpw6f5qg-8000.brs.devtunnels.ms"}
		allowedMethods := []string{"GET", "POST", "OPTIONS"}
		allowedHeaders := []string{"Content-Type"}

		// Verificar si el origen de la solicitud está permitido
		origin := r.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				// Establecer los encabezados CORS
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ","))
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowedHeaders, ","))

				// Continuar con la siguiente middleware o handler
				next.ServeHTTP(w, r)
				return
			}
		}

		// Si el origen no está permitido, responder con un error
		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}
