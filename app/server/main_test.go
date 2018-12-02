package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/", strings.NewReader(""))

	mux.ServeHTTP(writer, request)

	if code := writer.Code; code != 200 {
		t.Errorf("response code = %v != 200", code)
	}
}
