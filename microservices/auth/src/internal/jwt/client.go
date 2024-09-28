package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type Client struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWTClient() *Client {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatalf("error generating random private token: %v", err)
	}
	err = key.Validate()
	if err != nil {
		log.Fatalf("error validating private key: %v", err)
	}

	return &Client{
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
