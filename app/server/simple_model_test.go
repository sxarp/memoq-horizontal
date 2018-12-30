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
