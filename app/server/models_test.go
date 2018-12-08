package main

import (
	"testing"
)

func TestBasicDBOperaions(t *testing.T) {
	inputValue := "hoge"

	if res := UpdateEntityProgressive(inputValue); res != inputValue {
		t.Errorf("Expected %s, got %s", inputValue, res)
	}
}
