package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func reqRes(srv *server, method, path string, body io.Reader) *httptest.ResponseRecorder {

	writer := httptest.NewRecorder()

	request, _ := http.NewRequest(method, path, body)

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

func TestSimpleCreate(t *testing.T) {
	type input struct {
		Name string `json:"name"`
	}

	name := "test"

	mapBody := &input{
		Name: name,
	}

	body, _ := json.Marshal(mapBody)

	res := reqRes(genSrv(), "POST", "/simple/create", strings.NewReader(string(body)))

	fmt.Println(res.Code)

	fmt.Println(res.Body)

	if code := res.Code; code != 200 {
		t.Errorf("response code = %v != 200", code)
	}

	if body := res.Body.String(); body != name {
		t.Errorf("expected %v, got %v", name, body)
	}
}
