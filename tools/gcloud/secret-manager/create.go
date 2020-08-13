package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"github.com/google/subcommands"
)

type createCmd struct {
	projectID, key, value, path string
}

func newCreateCmd() *createCmd {
	return &createCmd{}
}

func (*createCmd) Name() string {
	return "create"
}

func (*createCmd) Synopsis() string {
	return "create secret"
}

func (*createCmd) Usage() string {
	return `usage: create secret`
}

func (cmd *createCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.projectID, "p", "", "project id")
	f.StringVar(&cmd.key, "k", "", "key")
	f.StringVar(&cmd.value, "v", "", "value")
	f.StringVar(&cmd.path, "f", "", "file path")
}

func (cmd *createCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if cmd.projectID == "" || cmd.key == "" || (cmd.value == "" && cmd.path == "") {
		log.Println("need -p [gcp project id] -k [secret key] -v [secret value] or -f [secret file path]")
		return subcommands.ExitFailure
	}

	var client *secretmanager.Client
	{
		var err error
		client, err = secretmanager.NewClient(ctx)
		if err != nil {
			log.Fatalf("failed to setup client: %v", err)
		}
	}

	// Create the request to create the secret.
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", cmd.projectID),
		SecretId: cmd.key,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}

	var secret *secretmanagerpb.Secret
	{
		var err error
		secret, err = client.CreateSecret(ctx, createSecretReq)
		if err != nil {
			log.Fatalf("failed to create secret: %v", err)
		}
	}

	// Declare the payload to storage.
	var payload []byte
	if cmd.value != "" {
		payload = []byte(cmd.value)
	}
	if cmd.path != "" {
		ba, err := ioutil.ReadFile(cmd.path)
		if err != nil {
			log.Fatalf("failed to read secret file: %+v", err)
		}
		payload = ba
	}
	if payload == nil {
		log.Fatal("payload is nil")
	}

	// Build the request.
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}

	var accessRequest *secretmanagerpb.AccessSecretVersionRequest
	{
		// Call the API.
		version, err := client.AddSecretVersion(ctx, addSecretVersionReq)
		if err != nil {
			log.Fatalf("failed to add secret version: %v", err)
		}

		// Build the request.
		accessRequest = &secretmanagerpb.AccessSecretVersionRequest{
			Name: version.Name,
		}
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}

	// Print the secret payload.
	//
	// WARNING: Do not print the secret in a production environment - this
	// snippet is showing how to access the secret material.
	log.Printf("Plaintext: %s", result.Payload.Data)

	return subcommands.ExitSuccess
}
