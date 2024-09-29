package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"nurlashko.dev/auth/internal"
	"nurlashko.dev/auth/internal/auth"
	"nurlashko.dev/auth/internal/handler"
	"nurlashko.dev/auth/internal/jwt"
	"nurlashko.dev/auth/internal/proxy"
)

type AuthApp struct {
	config internal.Config
}

func NewAuthApp() *AuthApp {
	config, err := internal.ParseConfig()
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return &AuthApp{
		config: config,
	}
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	mux := http.NewServeMux()
	app := NewAuthApp()
	jwtClient := jwt.NewJWTClient()
	vaultClient := auth.NewVaultClient(app.config)
	statikaProxy := proxy.NewStatikaReverseProxy()

	mux.HandleFunc("GET /public/jwt-key", handler.GetJWTPublicKey(jwtClient))
	mux.HandleFunc("POST /token", handler.SetCookieJWTToken(jwtClient, vaultClient))
	go statikaProxy.StartProxy()

	slog.Info("Listening on :8000")
	if err := http.ListenAndServe("0.0.0.0:8000", mux); err != nil {
		slog.Error("error listening: %v", err)
	}
}
