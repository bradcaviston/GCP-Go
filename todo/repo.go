package todo

import (
	"gcp-go/db"

	"cloud.google.com/go/firestore"
)

func repo()  *firestore.CollectionRef {
	return db.Client.Collection("todos")
}