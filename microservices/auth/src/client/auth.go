package client

import (
	"context"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

type AuthClient struct {
	vault *vault.Client
}

const prodAddress = "http://vault.auth:8200"
const debugAddress = "https://vault.nurlashko.dev"

func NewAuthClient(debug bool) *AuthClient {
	vaultAddress := prodAddress
	if debug {
		vaultAddress = debugAddress
	}
	client, err := vault.New(
		vault.WithAddress(vaultAddress),
		vault.WithRequestTimeout(5*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return &AuthClient{
		vault: client,
	}
}

func (c *AuthClient) IsTokenValid(token string) bool {
	ctx := context.Background()
	_, err := c.vault.Auth.TokenLookUpSelf(ctx, vault.WithToken(token))
	return err == nil
}

func (c *AuthClient) GetClientToken(ghToken string) (string, error) {
	ctx := context.Background()
	r, err := c.vault.Auth.GithubLogin(ctx, schema.GithubLoginRequest{Token: ghToken})
	if err != nil {
		return "", err
	}
	return r.Auth.ClientToken, nil
}
