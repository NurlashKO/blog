package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/NurlashKO/blog/microservices/auth/src/internal/auth"
	"github.com/NurlashKO/blog/microservices/auth/src/internal/jwt"
)

type publicKeyResponse struct {
	PublicKey []byte `json:"public_key"`
}

func SetCookieJWTToken(jwt *jwt.Client, auth *auth.VaultClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ghToken := r.FormValue("gh_token")
		user, err := auth.GetEntity(ghToken)
		if err != nil {
			http.Error(w, "failed to get auth entity", http.StatusUnauthorized)
			slog.Error("failed to get auth entity: " + err.Error())
			return
		}
		token, err := jwt.GenerateSignedClaim(user)
		if err != nil {
			http.Error(w, "failed to get jwt token", http.StatusUnauthorized)
			slog.Error("failed to get jwt token: " + err.Error())
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "X-AUTH-TOKEN",
			Value:    token,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Domain:   jwt.Domain,
		})
		w.Header().Set("X-AUTH-TOKEN", token)
		http.Redirect(w, r, "https://blog.nurlashko.dev", 301)
	}
}

func GetJWTPublicKey(jwt *jwt.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := publicKeyResponse{
			PublicKey: jwt.GetPublicKey(),
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(t)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			slog.Error("failed to marshal response: " + err.Error())
			return
		}
		w.Write(data)
	}
}
