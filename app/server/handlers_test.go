package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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

	mapBody := map[string]string{
		"hoge": "fuga",
	}

	fmt.Println("run tests!")

	body, _ := json.Marshal(mapBody)

	res := reqRes(genSrv(), "POST", "/simple/create", bytes.NewReader(body))

	fmt.Println(res.Code)

	fmt.Println(res.Body)

	if code := res.Code; code != 200 {
		t.Errorf("response code = %v != 200", code)
	}
}
