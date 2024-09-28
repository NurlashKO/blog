package auth

import (
	"context"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"

	"nurlashko.dev/auth/internal"
)

type VaultClient struct {
	vault *vault.Client
}

const prodAddress = "http://vault:8200"
const debugAddress = "https://vault.nurlashko.dev"

func NewVaultClient(config internal.Config) *VaultClient {
	vaultAddress := prodAddress
	if config.Debug {
		vaultAddress = debugAddress
	}
	client, err := vault.New(
		vault.WithAddress(vaultAddress),
		vault.WithRequestTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return &VaultClient{
		vault: client,
	}
}

func (c *VaultClient) GetEntity(ghToken string) (string, error) {
	ctx := context.Background()
	r, err := c.vault.Auth.GithubLogin(ctx, schema.GithubLoginRequest{Token: ghToken})
	if err != nil {
		return "", err
	}
	return r.Auth.Metadata["username"], nil
}
