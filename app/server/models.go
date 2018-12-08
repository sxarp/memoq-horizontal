package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"time"
)

type Entity struct {
	Value string
}

func UpdateEntity(inputValue string) string {
	project := "my-project-id"
	keyToUpdate := "test_key"

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
	defer cancel() // releases resources if slowOperation completes before timeout elapses

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := datastore.NewClient(ctx, project)
	if err != nil {
		fmt.Printf("Failed to create DS client.")
		panic("failed")
	}

	k := datastore.NameKey("Entity", keyToUpdate, nil)

	eClear := &Entity{Value: ""}
	if _, err := dsClient.Put(ctx, k, eClear); err != nil {
		fmt.Printf("Failed to Put.")
		fmt.Printf(err.Error())
		panic("failed")
	}

	e := new(Entity)
	if err := dsClient.Get(ctx, k, e); err != nil {
		fmt.Printf("Failed to Get.")
		panic("failed")
	}

	old := e.Value
	e.Value = inputValue

	if _, err := dsClient.Put(ctx, k, e); err != nil {
		fmt.Printf("Failed to Put.")
		fmt.Printf(err.Error())
		panic("failed")
	}

	fmt.Printf("Updated value from %q to %q\n", old, e.Value)

	retEn := new(Entity)
	if err := dsClient.Get(ctx, k, retEn); err != nil {
		fmt.Printf("Failed to Get.")
		panic("failed")
	}

	return retEn.Value
}
