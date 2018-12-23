package main

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
)

type DStore struct {
	prj     string
	kind    string
	timeout int
	client  datastore.Client
}

func (d *DStore) ctxCan() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(d.timeout)*time.Millisecond)

	return ctx, cancel
}

func NewDStore(prj, kind string, timeout int) DStore {

	client, err := datastore.NewClient(context.Background(), prj)

	if err != nil {
		panic("Failed to create DataStore client.")
	}

	return DStore{prj: prj, timeout: timeout, client: *client, kind: kind}
}

func (d *DStore) key(key string) *datastore.Key {
	return datastore.NameKey(d.kind, key, nil)
}

func (d *DStore) Get(key string, container interface{}) error {
	ctx, cancel := d.ctxCan()
	defer cancel()

	return d.client.Get(ctx, d.key(key), container)
}

func (d *DStore) Put(key string, container interface{}) (interface{}, error) {
	ctx, cancel := d.ctxCan()
	defer cancel()

	ret, err := d.client.Put(ctx, d.key(key), container)

	return ret, err
}

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
