package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func reqRes(srv *server, method, path, body string) *httptest.ResponseRecorder {

	writer := httptest.NewRecorder()

	request, err := http.NewRequest(method, path, strings.NewReader(body))

	if err != nil {
		panic(err)
	}

	srv.ServeHTTP(writer, request)

	return writer
}

func genSrv() (srv *server) {
	srv = &server{
		router: mux.NewRouter(),
	}

	srv.routes()

	return
}

func genJson(str interface{}) string {
	body, _ := json.Marshal(str)

	return string(body)
}

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
