package samlidp

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

// generateKeyPair generates a new RSA key pair and certificate
func generateKeyPair(domain string) (*rsa.PrivateKey, *x509.Certificate, error) {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	// Create certificate template
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate serial number: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: domain,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour), // 1 year
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{domain},
	}

	// Create certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create certificate: %v", err)
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse certificate: %v", err)
	}

	return privateKey, cert, nil
}

// saveKeyPair saves the private key and certificate to files
func saveKeyPair(keyFile, certFile string, privateKey *rsa.PrivateKey, cert *x509.Certificate) error {
	// Save private key
	keyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if err := os.WriteFile(keyFile, pem.EncodeToMemory(keyPEM), 0600); err != nil {
		return fmt.Errorf("failed to write private key: %v", err)
	}

	// Save certificate
	certPEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	}
	if err := os.WriteFile(certFile, pem.EncodeToMemory(certPEM), 0644); err != nil {
		return fmt.Errorf("failed to write certificate: %v", err)
	}

	return nil
}

// ensureKeyPair ensures that the key pair exists, generating it if necessary
func ensureKeyPair(keyFile, certFile, domain string) (*rsa.PrivateKey, *x509.Certificate, error) {
	// Check if files exist
	_, keyErr := os.Stat(keyFile)
	_, certErr := os.Stat(certFile)

	// If both files exist, load them
	if keyErr == nil && certErr == nil {
		return loadKeyPair(keyFile, certFile)
	}

	// Generate new key pair
	privateKey, cert, err := generateKeyPair(domain)
	if err != nil {
		return nil, nil, err
	}

	// Save key pair
	if err := saveKeyPair(keyFile, certFile, privateKey, cert); err != nil {
		return nil, nil, err
	}

	return privateKey, cert, nil
}

// loadKeyPair loads the key pair from files
func loadKeyPair(keyFile, certFile string) (*rsa.PrivateKey, *x509.Certificate, error) {
	keyBytes, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read key file: %v", err)
	}

	certBytes, err := os.ReadFile(certFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read cert file: %v", err)
	}

	keyBlock, _ := pem.Decode(keyBytes)
	if keyBlock == nil {
		return nil, nil, fmt.Errorf("failed to decode private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	certBlock, _ := pem.Decode(certBytes)
	if certBlock == nil {
		return nil, nil, fmt.Errorf("failed to decode certificate")
	}

	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse certificate: %v", err)
	}

	return key, cert, nil
}
