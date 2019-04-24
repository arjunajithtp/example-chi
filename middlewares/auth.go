package middlewares

import (
	"net/http"
	"strings"
)

// CustomAuth is for our custom authentication
func CustomAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		tokenPieces := strings.Split(token, " ")
		if len(tokenPieces) == 2 && tokenPieces[0] == "Basic" {
			// TODO: do basic auth
			handler.ServeHTTP(w, r)
			return
		}
		// TODO: do session auth

		// if both validations fails
		http.Error(w, "error while trying to authenticate", http.StatusUnauthorized)

	})
}
