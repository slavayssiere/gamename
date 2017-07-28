package common

import (
	"os"

	"log"

	vault "github.com/hashicorp/vault/api"
)

// VaultClient pointer to a client vault
type VaultClient struct {
	vault *vault.Client
}

//VaultManagement returns a Client interface for given consul address
func VaultManagement() (*VaultClient, error) {
	config := vault.DefaultConfig()
	addr := os.Getenv("VAULT_HOST")
	if len(addr) == 0 {
		addr = "http://127.0.0.1:8200"
	}
	config.Address = addr

	c, err := vault.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &VaultClient{vault: c}, nil
}

// GetSecret get secrets
func GetSecret(client *VaultClient, path string) map[string]interface{} {
	secret, err := client.vault.Logical().Read(path)
	if err != nil {
		log.Fatalln(err)
	}
	return secret.Data
}
