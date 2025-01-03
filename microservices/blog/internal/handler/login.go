package handler

import (
	"log/slog"
	"net/http"

	auth "nurlashko.dev/auth/client"
	"nurlashko.dev/blog/internal/view/user"
)

func LoginGET() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := user.LoginModal().Render(r.Context(), w)
		if err != nil {
			slog.Error("error rendering: %v", err)
		}
	}
}

func LoginPOST(auth *auth.AuthClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ghToken := r.FormValue("gh_token")
		token, err := auth.GetClientToken(ghToken)
		if err != nil {
			http.Error(w, "failed to get token", http.StatusUnauthorized)
			slog.Error("failed to get token: " + err.Error())
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "X-AUTH-TOKEN",
			Value:    token,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Domain:   "nurlashko.dev",
		})
		w.WriteHeader(http.StatusOK)
	}
}
