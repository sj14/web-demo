package gcloud

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

func NewGCloudStore() GCloudStore {
	myGcloudStore := GCloudStore{}

	ctx := context.Background()

	// Creates a client.
	var err error
	myGcloudStore.client, err = storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return myGcloudStore
}

type GCloudStore struct {
	client *storage.Client
}
