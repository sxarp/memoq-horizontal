package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var errResp = map[string]string{"status": "error"}

func (s *server) simpleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var ret interface{} = nil

		defer func() {
			js, _ := json.Marshal(ret)
			fmt.Fprintf(w, string(js))
		}()

		r := &Request{req}

		e := &Simple{}

		if err := r.JsonBody(e); err != nil {
			w.WriteHeader(400)
			ret = errResp
			return
		}

		id, err := e.Save(Repo(s.repo))

		if err != nil {
			w.WriteHeader(500)
			ret = errResp
			return
		}

		ret = map[string]int{"id": id}
	}
}
