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
	timeout := 100

	d := NewDStore(prj, timeout)

	if d.prj != prj {
		t.Errorf("Expected %s, got %s", prj, d.prj)
	}

	if d.timeout != timeout {
		t.Errorf("Expected %d, got %d", timeout, d.timeout)
	}
}
