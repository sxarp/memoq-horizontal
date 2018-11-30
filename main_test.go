package main

import "testing"

func TestAdder(t *testing.T) {
	if res := Adder(1, 2); res != 3 {
		t.Error("wrong!")
	}
}
