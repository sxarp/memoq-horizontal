package main

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
)

func RefreshDStore(d *DStore) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	client, _ := datastore.NewClient(context.Background(), d.prj)

	q := datastore.NewQuery(d.kind).KeysOnly()

	if keys, err := client.GetAll(ctx, q, nil); err != nil {
		panic("Failed to GetAll.")

	} else {
		client.DeleteMulti(ctx, keys)
	}

}

func TestCtxCan(t *testing.T) {
	d := DStore{timeout: 500}

	ctx, can := d.ctxCan()

	if ctx == nil || can == nil {
		t.Errorf("#ctxCan produced wrong value.")
	}
}

func TestNewDStore(t *testing.T) {

	prj := "test"
	kind := "kind"
	timeout := 500

	d := NewDStore(prj, kind, timeout)

	if d.prj != prj {
		t.Errorf("Expected %s, got %s.", prj, d.prj)
	}

	if d.timeout != timeout {
		t.Errorf("Expected %d, got %d.", timeout, d.timeout)
	}
}

func TestPutGet(t *testing.T) {
	prj := "test"
	kind := "kind"
	timeout := 500

	d := NewDStore(prj, kind, timeout)
	defer RefreshDStore(&d)

	type Ent struct {
		Name string
		Age  int
	}

	cont := &Ent{}
	if d.Get(d.NameKey("foo"), cont); cont.Name != "" || cont.Age != 0 {
		t.Errorf("Expected null values, got %+v.", cont)

	}

	key := "kyutie"
	age := 360
	name := "Sophie"

	cont = &Ent{Age: age, Name: name}

	if _, err := d.Put(d.NameKey(key), cont); err != nil {
		t.Errorf("Failed to put: %s", err)

	}
	if cont.Name != name || cont.Age != age {
		t.Errorf("Expected the same values, got %+v.", cont)

	}

	cont = &Ent{}
	if d.Get(d.NameKey(key), cont); cont.Name != name || cont.Age != age {
		t.Errorf("Expected the putted values, got %+v.", cont)

	}

}

func TestCreateDelete(t *testing.T) {
	prj := "test"
	kind := "kind"
	timeout := 500

	d := NewDStore(prj, kind, timeout)
	defer RefreshDStore(&d)

	type Ent struct {
		Name string
		Age  int
	}

	age := 450
	name := "Ellie"
	cont := &Ent{Age: age, Name: name}

	key, err := d.Create(cont)

	if err != nil {
		t.Errorf("Got error: %+v.", err)

	}

	cont = &Ent{}
	if d.Get(key, cont); cont.Name != name || cont.Age != age {
		t.Errorf("Expected the putted values, got %+v.", cont)

	}

	if err := d.Delete(key); err != nil {
		t.Errorf("Failed to delete: %s.", err)
	}

	cont = &Ent{}
	if d.Get(key, cont); cont.Name != "" || cont.Age != 0 {
		t.Errorf("Expected null values, got %+v.", cont)

	}

}

func TestCheckKey(t *testing.T) {
	prj := "test"
	timeout := 500

	dA := NewDStore(prj, "kindA", timeout)
	defer RefreshDStore(&dA)
	dB := NewDStore(prj, "kindB", timeout)
	defer RefreshDStore(&dB)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to raise panic, got nothing.")
		}
	}()

	dA.checkKey(dB.NameKey("hoge"))
}
