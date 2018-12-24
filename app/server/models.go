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

func (d *DStore) NameKey(key string) *datastore.Key {
	return datastore.NameKey(d.kind, key, nil)
}

func (d *DStore) checkKey(key *datastore.Key) {
	if key.Kind != d.kind {
		panic(fmt.Sprintf("Wrong kind of key specified, got %s, expected %s.", key.Kind, d.kind))

	}
}

//func (d *DStore) IDKey(id int64) *datastore.Key {
//return datastore.IDKey(d.kind, id, nil)
//}

func (d *DStore) Get(key *datastore.Key, container interface{}) error {
	d.checkKey(key)

	ctx, cancel := d.ctxCan()
	defer cancel()

	return d.client.Get(ctx, key, container)
}

func (d *DStore) Put(key *datastore.Key, container interface{}) (*datastore.Key, error) {
	d.checkKey(key)

	ctx, cancel := d.ctxCan()
	defer cancel()

	key, err := d.client.Put(ctx, key, container)

	return key, err
}

func (d *DStore) Create(container interface{}) (*datastore.Key, error) {
	ctx, cancel := d.ctxCan()
	defer cancel()

	newKey := datastore.IncompleteKey(d.kind, nil)

	key, err := d.client.Put(ctx, newKey, container)

	return key, err
}

func (d *DStore) Delete(key *datastore.Key) error {
	d.checkKey(key)

	ctx, cancel := d.ctxCan()
	defer cancel()

	return d.client.Delete(ctx, key)
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
