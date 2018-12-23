package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func reqRes(srv *server, method, path, body string) *httptest.ResponseRecorder {

	writer := httptest.NewRecorder()

	request, _ := http.NewRequest(method, path, strings.NewReader(body))

	srv.ServeHTTP(writer, request)

	return writer
}

func genSrv() (srv *server) {
	srv = &server{
		router: http.NewServeMux(),
	}

	srv.routes()

	return
}

func TestSimpleCreate(t *testing.T) {

	res := reqRes(genSrv(), "GET", "/", "")

	if code := res.Code; code != 200 {
		t.Errorf("response code = %v != 200", code)
	}
}
