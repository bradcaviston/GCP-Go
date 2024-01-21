package todo

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var ctx = context.Background()

func findAll() ([]Todo, error) {
	iter := repo().Documents(ctx)
	var todos []Todo
	var returnError error

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to iterate: %v", err)
			returnError = err
		}

		var todo Todo
		doc.DataTo(&todo)
		todo.Id = doc.Ref.ID

		todos = append(todos, todo)
	}

	return todos, returnError
}

func findOne(id string) (Todo, error) {
	var todo Todo

	doc, err := repo().Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Failed finding todo: %v", err)
		return todo, err
	}

	
	err = doc.DataTo(&todo)
	todo.Id = doc.Ref.ID

	return todo, err
}

func upsert(todo Todo) (Todo, error) {
	id := todo.Id
	var docRef *firestore.DocumentRef

	if (id != "") {
		docRef = repo().Doc(id)
		_, err := docRef.Set(ctx, todo)
		if err != nil {
			log.Printf("Failed updating todo: %v", err)
		}
	} else {
		addRef, _, err := repo().Add(ctx, todo)
		docRef = addRef
		if err != nil {
			log.Printf("Failed adding todo: %v", err)
		}
	}

	doc, err := docRef.Get(ctx)
	if err != nil {
		log.Printf("Failed to write todo: %v", err)
	}

	doc.DataTo(&todo)
	todo.Id = docRef.ID

	return todo, err
}

func delete(id string) error {
	_, err := repo().Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Failed to delete todo: %v", err)
	}

	return err
}
