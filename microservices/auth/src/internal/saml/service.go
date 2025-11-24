package samlidp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/crewjam/saml"
	"github.com/crewjam/saml/logger"
	"github.com/NurlashKO/blog/microservices/auth/src/internal"
	"github.com/NurlashKO/blog/microservices/auth/src/internal/jwt"
)

// Service represents the SAML Identity Provider service
type Service struct {
	config *internal.Config
	idp    *saml.IdentityProvider
	logger logger.Interface
	jwt    *jwt.Client

	// Mutex to protect concurrent access to IDP configuration
	idpConfigMu sync.RWMutex
}

// NewService creates a new SAML Identity Provider service
func NewService(config *internal.Config, jwtClient *jwt.Client) (*Service, error) {
	// Ensure key pair exists
	privateKey, cert, err := ensureKeyPair(config.SAML.KeyFile, config.SAML.CertFile, config.Domain)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure key pair: %w", err)
	}

	// Parse URL
	baseURL, err := url.Parse(config.SAML.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid SAML URL: %w", err)
	}

	// Ensure URL path ends without slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	// Construct metadata and SSO URLs
	metadataURL := *baseURL
	metadataURL.Path += "/metadata"
	ssoURL := *baseURL
	ssoURL.Path += "/sso"

	// Create service
	service := &Service{
		config: config,
		logger: logger.DefaultLogger,
		jwt:    jwtClient,
	}

	// Create service provider
	sp, err := NewServiceProvider(config.SAML.ServiceProviders[0])
	if err != nil {
		return nil, fmt.Errorf("failed to create service provider: %w", err)
	}

	// Create Identity Provider
	service.idp = &saml.IdentityProvider{
		Key:                     privateKey,
		Signer:                  privateKey,
		Logger:                  service.logger,
		Certificate:             cert,
		MetadataURL:             metadataURL,
		SSOURL:                  ssoURL,
		LoginURL:                url.URL{Path: "/"}, // Use the root path for login
		ServiceProviderProvider: sp,
	}

	// Set up session provider
	service.idp.SessionProvider = NewSessionProvider(jwtClient, service.logger)

	return service, nil
}

// MetadataHandler returns the SAML metadata handler
func (s *Service) MetadataHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.idpConfigMu.RLock()
		defer s.idpConfigMu.RUnlock()
		s.idp.ServeMetadata(w, r)
	})
}

// SSOHandler returns the SAML Single Sign-On handler
func (s *Service) SSOHandler() http.Handler {
	return http.HandlerFunc(s.idp.ServeSSO)
}
