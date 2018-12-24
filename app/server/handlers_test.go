package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

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
