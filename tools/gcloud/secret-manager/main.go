package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func main() {
	flag.Parse()
	projectID := flag.Arg(0)
	if projectID == "" {
		log.Println("no projectID")
		return
	}

	key := flag.Arg(1)
	if key == "" {
		log.Println("no key")
		return
	}

	val := flag.Arg(2)
	if val == "" {
		log.Println("no value")
		return
	}

	ctx := context.Background()

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
		Parent:   fmt.Sprintf("projects/%s", projectID),
		SecretId: key,
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

	// Declare the payload to store.
	payload := []byte(val)

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
}
