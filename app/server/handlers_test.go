package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func genSrv(repo Repo) (srv *server) {
	srv = &server{
		router: mux.NewRouter(),
		repo:   repo,
		logger: log.New(ioutil.Discard, "logger: ", 0),
		// For debugging.
		// logger: log.New(os.Stdout, "logger: ", log.Llongfile|log.Lmicroseconds),
	}

	srv.routes()

	return
}

func genJson(str interface{}) string {
	body, _ := json.Marshal(str)

	return string(body)
}
