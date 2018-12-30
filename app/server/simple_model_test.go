package main

import (
	"reflect"
	"testing"
)

func TestSimpleModelSaveFindDestroy(t *testing.T) {

	d := NewDStore("test", 500)

	name := "Akari"
	age := 16

	s := Simple{Name: name, Age: age}
	s.SetKind(d)
	defer RefreshDStore(d)

	key, err := s.Save(d)

	if err != nil {
		t.Errorf("Failed to save: %v.", err)
	}

	ss := Simple{}

	ss.Find(d, key)

	if !reflect.DeepEqual(s, ss) {
		t.Errorf("Expected %v, got %v.", s, ss)
	}

	(&Simple{}).Destroy(d, key)

	ss = Simple{}

	if err := ss.Find(d, key); err == nil {
		t.Errorf("Expecred error, nothing raised.")
	} else if !reflect.DeepEqual(ss, Simple{}) {
		t.Errorf("Expected empty struct, got %v.", ss)

	}
}

func TestSimpleModelAllWithLimit(t *testing.T) {
	d := NewDStore("test", 500)

	name := "Hinata"

	for i := 0; i < 20; i++ {
		(&Simple{Name: name, Age: i}).Save(d)

	}

	fetchNum := 5

	s := []Simple{}
	ss := Simples(s)

	if err := (&ss).AllWithLimit(d, fetchNum); err != nil {
		t.Errorf("Failed to exec AllWithLimit: %s.", err)
	}

	if length := len(ss); length != fetchNum {
		t.Errorf("Expected length is %d, got %d.", fetchNum, length)
	}

	for i := 0; i < fetchNum; i++ {
		if reflect.DeepEqual(Simple{}, ss[i]) {
			t.Errorf("Got unexpected zero values: %v.", ss)
		}

	}

	(&Simple{}).SetKind(d)
	RefreshDStore(d)
}
