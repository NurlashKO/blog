package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
)

type Client struct {
	Domain     string
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWTClient(domain string) *Client {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatalf("error generating random private token: %v", err)
	}
	err = key.Validate()
	if err != nil {
		log.Fatalf("error validating private key: %v", err)
	}

	return &Client{
		Domain:     domain,
		privateKey: key,
		publicKey:  &key.PublicKey,
	}
}

func (s *Client) GenerateSignedClaim(user string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user": user,
	})
	signedClaim, err := t.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}
	return signedClaim, nil
}

func (s *Client) GetPublicKey() []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(s.publicKey),
	})
}

func (s *Client) VerifyToken(token string) bool {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodRS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return s.publicKey, nil
	})
	if err != nil {
		slog.Error("failed to parse token: %v", err)
		return false
	}
	return t.Valid
}

// ParseToken parses a JWT token and returns the token object
func (s *Client) ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodRS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return s.publicKey, nil
	})
}
