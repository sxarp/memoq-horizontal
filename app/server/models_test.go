package main

import (
	"testing"
)

func TestBasicDBOperaions(t *testing.T) {
	inputValue := "hoge"

	if res := UpdateEntity(inputValue); res != inputValue {
		t.Errorf("Expected %s, got %s.", inputValue, res)
	}
}

func TestCtxCan(t *testing.T) {
	d := DStore{timeout: 100}

	ctx, can := d.ctxCan()

	if ctx == nil || can == nil {
		t.Errorf("#ctxCan produced wrong value.")
	}
}

func TestNewDStore(t *testing.T) {

	prj := "test"
	kind := "kind"
	timeout := 100

	d := NewDStore(prj, kind, timeout)

	if d.prj != prj {
		t.Errorf("Expected %s, got %s", prj, d.prj)
	}

	if d.timeout != timeout {
		t.Errorf("Expected %d, got %d", timeout, d.timeout)
	}
}

func TestPutGet(t *testing.T) {
	prj := "test"
	kind := "kind"
	timeout := 100

	d := NewDStore(prj, kind, timeout)

	type Ent struct {
		Name string
		Age  int
	}

	cont := &Ent{}
	if d.Get(d.NameKey("foo"), cont); cont.Name != "" || cont.Age != 0 {
		t.Errorf("Expected null values, got %+v", cont)

	}

	key := "kyutie"
	age := 360
	name := "Sophie"

	cont = &Ent{Age: age, Name: name}
	if d.Put(d.NameKey(key), cont); cont.Name != name || cont.Age != age {
		t.Errorf("Expected the same values, got %+v", cont)

	}

	cont = &Ent{}
	if d.Get(d.NameKey(key), cont); cont.Name != name || cont.Age != age {
		t.Errorf("Expected the putted values, got %+v", cont)

	}
}

func TestCreateDelete(t *testing.T) {
	prj := "test"
	kind := "kind"
	timeout := 100

	d := NewDStore(prj, kind, timeout)

	type Ent struct {
		Name string
		Age  int
	}

	age := 450
	name := "Ellie"
	cont := &Ent{Age: age, Name: name}

	key, err := d.Create(cont)

	if err != nil {
		t.Errorf("Got error : %+v", err)

	}

	cont = &Ent{}
	if d.Get(key, cont); cont.Name != name || cont.Age != age {
		t.Errorf("expected the putted values, got %+v", cont)

	}

	d.Delete(key)

	cont = &Ent{}
	if d.Get(key, cont); cont.Name != "" || cont.Age != 0 {
		t.Errorf("expected null values, got %+v", cont)

	}

}

func TestCheckKey(t *testing.T) {
	prj := "test"
	timeout := 100

	dA := NewDStore(prj, "kindA", timeout)
	dB := NewDStore(prj, "kindB", timeout)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to raise panic, got nothing.")
		}
	}()

	dA.checkKey(dB.NameKey("hoge"))
}
