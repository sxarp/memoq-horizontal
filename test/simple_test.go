package main

import "testing"
import "app/simple"

func TestAdder(t *testing.T) {
	if res := simple.Adder(1, 2); res != 3 {
		t.Error("wrong!")
	}
}
