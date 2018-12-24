package main

import (
	"testing"
)

func TestSimpleCreate(t *testing.T) {
	type input struct {
		Name string `json:"name"`
	}

	name := "test"

	inputBody := &input{
		Name: name,
	}

	body := genJson(inputBody)

	res := reqRes(genSrv(), "POST", "/simple/create", body)

	if code := res.Code; code != 200 {
		t.Errorf("response code = %v != 200", code)
	}

	if body := res.Body.String(); body != name {
		t.Errorf("expected %v, got %v", name, body)
	}
}
