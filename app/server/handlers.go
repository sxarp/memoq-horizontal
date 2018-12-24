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
