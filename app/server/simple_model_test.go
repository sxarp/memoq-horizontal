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

	id, err := s.Save(d)

	if err != nil {
		t.Errorf("Failed to save: %v.", err)
	}

	ss := Simple{}

	ss.Find(d, id)

	if !reflect.DeepEqual(s, ss) {
		t.Errorf("Expected %v, got %v.", s, ss)
	}

	(&Simple{}).Destroy(d, id)

	ss = Simple{}

	if err := ss.Find(d, id); err == nil {
		t.Errorf("Expecred error, nothing raised.")
	} else if !reflect.DeepEqual(ss, Simple{}) {
		t.Errorf("Expected empty struct, got %v.", ss)

	}
}

func TestSimpleModelAllWithLimit(t *testing.T) {
	d := NewDStore("test", 500)
	(&Simple{}).SetKind(d)
	RefreshDStore(d)

	name := "Hinata"

	createNum := 21

	for i := 0; i < createNum; i++ {
		(&Simple{Name: name, Age: i}).Save(d)

	}

	s := []Simple{}
	ss := Simples(s)

	if err := (&ss).AllWithLimit(d, createNum); err != nil {
		t.Errorf("Failed to exec AllWithLimit: %s.", err)
	}

	if length := len(ss); length != createNum {
		t.Errorf("Expected length is %d, got %d.", createNum, length)
	}

	for i := 0; i < createNum; i++ {
		if expected := (Simple{Name: name, Age: i}); !reflect.DeepEqual(expected, ss[i]) {
			t.Errorf("Expected %v, got %v.", expected, ss[i])
		}

	}

	// The case: fetchNum > createNum.
	fetchNum := 33
	s = []Simple{}
	ss = Simples(s)
	if err := (&ss).AllWithLimit(d, fetchNum); err != nil {
		t.Errorf("Failed to exec AllWithLimit: %s.", err)
	}

	if length := len(ss); length != createNum {
		t.Errorf("Expected length is %d, got %d.", createNum, length)
	}
}
