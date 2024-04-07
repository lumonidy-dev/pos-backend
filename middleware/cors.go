package middleware

import "net/http"

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir cualquier origen
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Permitir cualquier método
		w.Header().Set("Access-Control-Allow-Methods", "*")
		// Permitir cualquier encabezado
		w.Header().Set("Access-Control-Allow-Headers", "*")
		// Continuar con el siguiente middleware
		next.ServeHTTP(w, r)
	})
}
