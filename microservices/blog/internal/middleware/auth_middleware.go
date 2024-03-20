package middleware

import (
	"net/http"

	"nurlashko.dev/blog/internal/client"
)

func AuthenticationMiddleware(auth *client.AuthClient, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("x-auth-token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !auth.IsTokenValid(cookie.Value) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
