package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func render(w http.ResponseWriter, status *int, resp *interface{}) {
	w.WriteHeader(*status)
	js, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(js))
}

var errResp = map[string]string{"status": "error"}
var okResp = map[string]string{"status": "ok"}

func (s *server) simpleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		status := 500
		var resp interface{} = errResp
		defer render(w, &status, &resp)

		r := &Request{req}

		sim := &Simple{}

		if err := r.JsonBody(sim); err != nil {
			status = 400
			resp = errResp
			return
		}

		id, err := sim.Save(Repo(s.repo))

		if err != nil {
			status = 500
			resp = errResp
			return
		}

		status = 200
		resp = map[string]int{"id": id}
	}
}

func (s *server) simpleShow() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		status := 500
		var resp interface{} = errResp
		defer render(w, &status, &resp)

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
		resp = sim
	}
}

func (s *server) simpleDestroy() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		status := 500
		var resp interface{} = errResp
		defer render(w, &status, &resp)

		id, convErr := strconv.Atoi(mux.Vars(req)["id"])

		if convErr != nil {
			status = 400
			return
		}

		if err := (&Simple{}).Find(s.repo, id); err != nil {
			if err.Error() == "datastore: no such entity" {
				// No content
				status = 204
				resp = okResp
				return

			}
			return
		}

		if err := (&Simple{}).Destroy(s.repo, id); err != nil {
			return
		}

		status = 200
		resp = okResp
	}
}

func (s *server) simpleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		status := 500
		var resp interface{} = errResp
		defer render(w, &status, &resp)

		limit := 10 // Default limit.

		if lims := req.FormValue("limit"); lims != "" {
			var convErr error = nil
			limit, convErr = strconv.Atoi(lims)
			if convErr != nil {
				status = 400
				return
			}
		}

		sim := []Simple{}
		sims := Simples(sim)

		if err := (&sims).AllWithLimit(s.repo, limit); err != nil {
			return
		}

		status = 200
		resp = sims
	}
}
