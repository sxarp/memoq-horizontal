package main

import (
	"reflect"
	"testing"
)

func TestSimpleModelSaveFind(t *testing.T) {

	ds := NewDStore("test", 500)
	d := &ds

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
}
