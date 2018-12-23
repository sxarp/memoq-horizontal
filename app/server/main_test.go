package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	srv := server{
		router: http.NewServeMux(),
	}

	srv.routes()

	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/test/index", strings.NewReader(""))

	srv.ServeHTTP(writer, request)

	if code := writer.Code; code != 200 {
		t.Errorf("response code = %v != 200", code)
	}
}
