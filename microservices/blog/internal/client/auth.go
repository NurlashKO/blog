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

func NewAuthClient() *AuthClient {
	client, err := vault.New(
		vault.WithAddress("https://vault.nurlashko.dev"),
		vault.WithRequestTimeout(3*time.Second),
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
