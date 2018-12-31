package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func render(w http.ResponseWriter, status *int, ret *interface{}) {
	w.WriteHeader(*status)
	js, _ := json.Marshal(ret)
	fmt.Fprintf(w, string(js))
}

var errResp = map[string]string{"status": "error"}
var okResp = map[string]string{"status": "ok"}

func (s *server) simpleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		status := 500
		var ret interface{} = errResp
		defer render(w, &status, &ret)

		r := &Request{req}

		sim := &Simple{}

		if err := r.JsonBody(sim); err != nil {
			status = 400
			ret = errResp
			return
		}

		id, err := sim.Save(Repo(s.repo))

		if err != nil {
			status = 500
			ret = errResp
			return
		}

		status = 200
		ret = map[string]int{"id": id}
	}
}

func (s *server) simpleShow() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		status := 500
		var ret interface{} = errResp
		defer render(w, &status, &ret)

		id, convErr := strconv.Atoi(mux.Vars(req)["id"])

		if convErr != nil {
			status = 400
			return
		}

		sim := Simple{}
		err := sim.Find(s.repo, id)

		if err != nil {
			if err.Error() == "datastore: no such entity" {
				status = 400
				return

			}
			return
		}

		status = 200
		ret = sim
	}
}

func (s *server) simpleDestroy() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		status := 500
		var ret interface{} = errResp
		defer render(w, &status, &ret)

		id, convErr := strconv.Atoi(mux.Vars(req)["id"])

		if convErr != nil {
			status = 400
			return
		}

		if err := (&Simple{}).Find(s.repo, id); err != nil {
			if err.Error() == "datastore: no such entity" {
				// No content
				status = 204
				ret = okResp
				return

			}
			return
		}

		if err := (&Simple{}).Destroy(s.repo, id); err != nil {
			return
		}

		status = 200
		ret = okResp
	}
}
