package main

import (
	"context"
	"log"
	"os"

	"github.com/leonardobiffi/vanta-sdk-go"
)

func main() {
	ctx := context.Background()

	v, err := vanta.New(
		ctx,
		vanta.WithOAuthCredentials(os.Getenv("VANTA_OAUTH_CLIENT_ID"), os.Getenv("VANTA_OAUTH_CLIENT_SECRET")),
		vanta.WithScopes(vanta.ScopeAllRead, vanta.ScopeAllWrite),
	)
	if err != nil {
		log.Fatalf("failed to initialize vanta sdk: %v", err)
	}

	err = v.UpdateResource(
		ctx,
		"68cb13bc4b91b274a5714dc0",
		vanta.UpdateResourceInput{
			InScope: false,
		},
	)
	if err != nil {
		log.Fatal("failed to update resource with vanta sdk", err)
	}

	log.Println("Resource updated successfully")
}
