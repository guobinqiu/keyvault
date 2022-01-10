package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"os"
	"strings"
)

type AzureClient struct {
	keyvaultClient keyvault.BaseClient
	vaultBaseURL   string
}

func NewAzureClient(vaultBaseURL string) (*AzureClient, error) {
	if isBlank(os.Getenv("AZURE_TENANT_ID")) {
		return nil, errors.New("missing environment variable AZURE_TENANT_ID")
	}
	if isBlank(os.Getenv("AZURE_CLIENT_ID")) {
		return nil, errors.New("missing environment variable AZURE_CLIENT_ID")
	}
	if isBlank(os.Getenv("AZURE_CLIENT_SECRET")) {
		return nil, errors.New("missing environment variable AZURE_CLIENT_SECRET")
	}

	authorizer, err := kvauth.NewAuthorizerFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("unable to create vault authorizer: %v", err)
	}

	basicClient := keyvault.New()
	basicClient.Authorizer = authorizer

	return &AzureClient{basicClient, vaultBaseURL}, nil
}

func (cli *AzureClient) GetValue(key string) (string, error) {
	secretResp, err := cli.keyvaultClient.GetSecret(context.Background(), cli.vaultBaseURL, key, "")
	if err != nil {
		return "", fmt.Errorf("unable to get value for secret: %v", err)
	}
	return *secretResp.Value, nil
}

func isBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
