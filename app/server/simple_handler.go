package main

import (
	"fmt"
	"net/http"
)

func (s *server) simpleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		r := &Request{req}

		type input struct {
			Name string `json:"name"`
		}

		var body input

		r.JsonBody(&body)

		fmt.Println(body.Name)

		fmt.Fprintf(w, body.Name)
	}
}
