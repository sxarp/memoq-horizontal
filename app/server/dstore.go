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

func NewDStore(prj string, timeout int) *DStore {

	client, err := datastore.NewClient(context.Background(), prj)

	if err != nil {
		panic("Failed to create DataStore client.")
	}

	return &DStore{prj: prj, timeout: timeout, client: *client}
}

func (d *DStore) SetKind(kind string) {
	d.kind = kind
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

func (d *DStore) NewQuery() *datastore.Query {
	return datastore.NewQuery(d.kind)
}

func (d *DStore) GetAll(q *datastore.Query, dst interface{}) ([]*datastore.Key, error) {
	ctx, cancel := d.ctxCan()
	defer cancel()

	keys, err := d.client.GetAll(ctx, q, dst)

	return keys, err
}
