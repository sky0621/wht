package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/api/iterator"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"github.com/google/subcommands"
)

type listCmd struct {
	projectID string
}

func newListCmd() *listCmd {
	return &listCmd{}
}

func (*listCmd) Name() string {
	return "list"
}

func (*listCmd) Synopsis() string {
	return "list secrets"
}

func (*listCmd) Usage() string {
	return `usage: list secrets`
}

func (cmd *listCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.projectID, "p", "", "project id")
}

func (cmd *listCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if cmd.projectID == "" {
		log.Println("need -p [gcp project id]")
		return subcommands.ExitFailure
	}

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Printf("failed to create secretmanager client: %v", err)
		return subcommands.ExitFailure
	}

	// Build the request.
	req := &secretmanagerpb.ListSecretsRequest{
		Parent: fmt.Sprintf("projects/%s", cmd.projectID),
	}

	// Call the API.
	it := client.ListSecrets(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Printf("failed to list secret versions: %v", err)
			return subcommands.ExitFailure
		}

		names := strings.Split(resp.Name, "/")
		reqName := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", names[1], names[3], "latest")

		// Build the request.
		req := &secretmanagerpb.AccessSecretVersionRequest{
			Name: reqName,
		}

		// Call the API.
		result, err := client.AccessSecretVersion(ctx, req)
		if err != nil {
			log.Printf("failed to access secret version: %v", err)
			return subcommands.ExitFailure
		}

		log.Printf("Found secret %s ... got value: %s\n", resp.Name, string(result.Payload.Data))
	}
	return subcommands.ExitSuccess
}
