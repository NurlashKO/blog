package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/NurlashKO/blog/microservices/auth/src/internal"
	"github.com/NurlashKO/blog/microservices/auth/src/internal/auth"
	"github.com/NurlashKO/blog/microservices/auth/src/internal/handler"
	"github.com/NurlashKO/blog/microservices/auth/src/internal/jwt"
	"github.com/NurlashKO/blog/microservices/auth/src/internal/proxy"
	samlidp "github.com/NurlashKO/blog/microservices/auth/src/internal/saml"
	"github.com/NurlashKO/blog/microservices/auth/src/internal/static"
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
	jwtClient := jwt.NewJWTClient(app.config.Domain)
	vaultClient := auth.NewVaultClient(app.config)

	statikaProxy := proxy.NewStatikaProxyTarget("https://static.nurlashko.dev", jwtClient)

	// Initialize SAML Identity Provider service
	samlService, err := samlidp.NewService(&app.config, jwtClient)
	if err != nil {
		log.Fatalf("error initializing SAML service: %v", err)
	}

	// SAML endpoints
	mux.Handle("/saml/metadata", samlService.MetadataHandler())
	mux.Handle("/saml/sso", samlService.SSOHandler())

	// General jwt auth endpoints
	mux.HandleFunc("GET /public/jwt-key", handler.GetJWTPublicKey(jwtClient))
	mux.HandleFunc("POST /token", handler.SetCookieJWTToken(jwtClient, vaultClient))

	// Default handler for unmatched routes
	mux.Handle("/", http.FileServerFS(static.GetPages()))

	go proxy.StartProxy(map[string]proxy.ProxyTarget{
		statikaProxy.Host: statikaProxy,
	})

	slog.Info("Listening on :8000")
	if err := http.ListenAndServe("0.0.0.0:8000", mux); err != nil {
		slog.Error("error listening: %v", err)
	}
}
