package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	*http.Request
}

func (r *Request) JsonBody(body interface{}) error {
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal([]byte(reqBody), &body); err != nil {
		return err
	}

	return nil
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi!")
	}
}

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
