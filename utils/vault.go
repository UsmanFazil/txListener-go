package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

const password string = ""

func GetDBPass() string {
	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	ctx := context.Background()

	secret, err := client.KVv2("secret").Get(ctx, "db-password")
	if err != nil {
		log.Fatalf(
			"Unable to read DB password from the vault: %v",
			err,
		)
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("Vault: data assertion failed")
	}

	return value

}
func GetPK() string {
	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	ctx := context.Background()

	secret, err := client.KVv2("secret").Get(ctx, "admin-pk")
	if err != nil {
		log.Fatalf(
			"Unable to read DB password from the vault: %v",
			err,
		)
	}

	value, ok := secret.Data["pk"].(string)
	if !ok {
		log.Fatalf("Vault: data assertion failed")
	}

	fmt.Println(value)
	return value

}
