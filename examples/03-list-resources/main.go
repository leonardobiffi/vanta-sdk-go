package main

import (
	"context"
	"fmt"
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

	opts := []vanta.ListResourcesOption{
		vanta.ResourceConnectionID("685062ce7a4eada920b6658b"),
		vanta.ResourceIsInScope(true),
		vanta.ResourceWithPageSize(10),
	}

	count := 0
	for {
		// List resources with the current options
		listResourcesOutput, err := v.ListResources(ctx, opts...)
		if err != nil {
			log.Fatal("failed to list resources with vanta sdk", err)
		}

		for _, resource := range listResourcesOutput.Results.Data {
			fmt.Printf("Resource: %s (%s) - %s\n", resource.DisplayName, resource.ResourceID, resource.Description)
			count++
		}

		// Check if there is a next page
		if listResourcesOutput.Results.PageInfo.HasNextPage {
			nextCursor := listResourcesOutput.Results.PageInfo.EndCursor
			opts = append(opts, vanta.ResourceWithPageCursor(nextCursor))
		} else {
			break
		}
	}

	fmt.Printf(">> Total resources listed: %d\n", count)
}
