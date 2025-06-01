package samlidp

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
	"nurlashko.dev/auth/internal"
)

// Service represents the SAML service
type Service struct {
	config *internal.Config
	sp     *samlsp.Middleware
}

// NewService creates a new SAML service
func NewService(config *internal.Config) (*Service, error) {
	// Ensure key pair exists
	privateKey, cert, err := ensureKeyPair(config.SAML.KeyFile, config.SAML.CertFile, config.Domain)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure key pair: %w", err)
	}

	// Parse URL
	url, err := url.Parse(config.SAML.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid SAML URL: %w", err)
	}

	// Create service provider
	sp, err := samlsp.New(samlsp.Options{
		URL:               *url,
		Key:               privateKey,
		Certificate:       cert,
		EntityID:          config.SAML.EntityID,
		AllowIDPInitiated: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create service provider: %w", err)
	}

	return &Service{
		config: config,
		sp:     sp,
	}, nil
}

// MetadataHandler returns the SAML metadata handler
func (s *Service) MetadataHandler() http.Handler {
	return http.HandlerFunc(s.sp.ServeMetadata)
}

// ACSHandler returns the SAML Assertion Consumer Service (SSO) handler
func (s *Service) ACSHandler() http.Handler {
	return http.HandlerFunc(s.sp.ServeACS)
}
