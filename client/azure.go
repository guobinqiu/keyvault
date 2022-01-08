package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"os"
)

type AzureClient struct {
	keyvaultClient keyvault.BaseClient
}

func NewAzureClient() (*AzureClient, error) {
	if os.Getenv("AZURE_TENANT_ID") == "" || os.Getenv("AZURE_CLIENT_ID") == "" || os.Getenv("AZURE_CLIENT_SECRET") == "" || os.Getenv("KVAULT") == "" {
		return nil, errors.New("env vars not set, exiting...")
	}

	authorizer, err := kvauth.NewAuthorizerFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("unable to create vault authorizer: %v", err)
	}

	basicClient := keyvault.New()
	basicClient.Authorizer = authorizer

	return &AzureClient{basicClient}, nil
}

func (cli *AzureClient) GetValue(key string) (string, error) {
	vaultName := os.Getenv("KVAULT")
	secretResp, err := cli.keyvaultClient.GetSecret(context.Background(), "https://"+vaultName+".vault.azure.cn", key, "")
	if err != nil {
		return "", fmt.Errorf("unable to get value for secret: %v", err)
	}
	return *secretResp.Value, nil
}
