package samlidp

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/crewjam/saml"
)

// ServiceProvider implements saml.ServiceProviderProvider interface
type ServiceProvider struct {
	metadata *saml.EntityDescriptor
}

// NewServiceProvider creates a new service provider from metadata URL
func NewServiceProvider(metadataURL string) (*ServiceProvider, error) {
	resp, err := http.Get(metadataURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch metadata: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch metadata: status %d", resp.StatusCode)
	}

	metadata := &saml.EntityDescriptor{}
	if err := xml.NewDecoder(resp.Body).Decode(metadata); err != nil {
		return nil, fmt.Errorf("failed to parse metadata: %w", err)
	}

	return &ServiceProvider{
		metadata: metadata,
	}, nil
}

// GetServiceProvider implements saml.ServiceProviderProvider
func (sp *ServiceProvider) GetServiceProvider(r *http.Request, serviceProviderID string) (*saml.EntityDescriptor, error) {
	if sp.metadata.EntityID != serviceProviderID {
		return nil, fmt.Errorf("unknown service provider: %s", serviceProviderID)
	}

	return sp.metadata, nil
}
