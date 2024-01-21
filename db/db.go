package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

var Client *firestore.Client

func Connect(ctx context.Context) {    
    client, err := firestore.NewClient(ctx, "test-326802")
    if err != nil {
        log.Fatal(err)
    }

    Client = client
}