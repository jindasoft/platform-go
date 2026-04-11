package xvault

import (
	"context"
	"fmt"
	"os"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
	"github.com/jindasoft/jinda-platforms/xlogger"
	"github.com/mitchellh/mapstructure"
)

func LoadVaultSecret[T any](vaultUrl, appName, kvName string, secret T) (*T, error) {
	// 1) Initialize Vault client/auth
	config := vault.DefaultConfig()

	config.Address = vaultUrl
	client, err := vault.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize Vault client: %w", err)
	}

	roleID := os.Getenv("APPROLE_ROLE_ID")
	if roleID == "" {
		return nil, fmt.Errorf("no role-id in APPROLE_ROLE_ID env var")
	}

	secretID := &auth.SecretID{FromEnv: "APPROLE_SECRET_ID"}
	appRoleAuth, err := auth.NewAppRoleAuth(roleID, secretID)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize AppRole auth: %w", err)
	}

	if _, err := client.Auth().Login(context.Background(), appRoleAuth); err != nil {
		return nil, fmt.Errorf("unable to login to AppRole: %w", err)
	}

	// 2) Read secret data from Vault
	rawSecret, err := client.KVv2(kvName).Get(context.Background(), appName)
	if err != nil {
		return nil, fmt.Errorf("unable to read secret: %w", err)
	}

	// 3) Docode secret to struct
	cfg := &mapstructure.DecoderConfig{
		Result:           &secret,
		WeaklyTypedInput: true,
	}

	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize mapstructure: %w", err)
	}

	if err := decoder.Decode(rawSecret.Data); err != nil {
		return nil, fmt.Errorf("unable to decode mapstructure: %w", err)
	}

	xlogger.SysInfof("Vault initialized")

	return &secret, nil
}
