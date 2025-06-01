package samlidp

import (
	"net/http"
	"time"

	"github.com/crewjam/saml"
	"github.com/crewjam/saml/logger"
	jwtlib "github.com/golang-jwt/jwt/v5"
	"nurlashko.dev/auth/internal/jwt"
)

// SessionProvider implements saml.SessionProvider interface
type SessionProvider struct {
	jwtClient *jwt.Client
	logger    logger.Interface
}

// NewSessionProvider creates a new session provider
func NewSessionProvider(jwtClient *jwt.Client, logger logger.Interface) *SessionProvider {
	return &SessionProvider{
		jwtClient: jwtClient,
		logger:    logger,
	}
}

// GetSession implements saml.SessionProvider
func (s *SessionProvider) GetSession(w http.ResponseWriter, r *http.Request, req *saml.IdpAuthnRequest) *saml.Session {
	// Get JWT token from cookie
	cookie, err := r.Cookie("X-AUTH-TOKEN")
	if err != nil {
		return nil
	}

	// Verify JWT token
	if !s.jwtClient.VerifyToken(cookie.Value) {
		return nil
	}

	// Parse JWT token to get user info
	token, err := s.jwtClient.ParseToken(cookie.Value)
	if err != nil {
		s.logger.Printf("failed to parse JWT token: %v", err)
		return nil
	}

	// Create SAML session
	return &saml.Session{
		ID:         token.Claims.(jwtlib.MapClaims)["user"].(string),
		CreateTime: time.Now(),
		ExpireTime: time.Now().Add(24 * time.Hour), // Session expires in 24 hours
		Index:      token.Claims.(jwtlib.MapClaims)["user"].(string),
		NameID:     token.Claims.(jwtlib.MapClaims)["user"].(string),
	}
}
