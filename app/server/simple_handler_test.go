package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"testing"
)

func TestSimpleHandlerCreate(t *testing.T) {

	d := NewDStore("test", 1000)
	(&Simple{}).SetKind(d)
	defer RefreshDStore(d)

	srv := genSrv(d)

	s := Simple{
		Name: "Sophie",
		Age:  360,
	}

	// Normal case.
	body := genJson(&s)

	res := reqRes(srv, "POST", "/simple", body)

	if code := res.Code; code != 200 {
		t.Errorf("Response code = %v != 200.", code)
	}

	resBody := res.Body.String()
	re := regexp.MustCompile("{\"id\":([0-9]+)}")
	if matched := re.FindStringSubmatch(resBody); len(matched) == 0 {
		t.Errorf("Failed to match expected pattern: %s.", resBody)
	} else if id, err := strconv.Atoi(matched[1]); err != nil {
		t.Errorf("Failed to convert id to integer.")
	} else if ss := (Simple{}); ss.Find(d, id) != nil {
		t.Errorf("Failed to find the saved value: %s.", err)
	} else if !reflect.DeepEqual(ss, s) {
		t.Errorf("Expected %v, got %v.", s, ss)
	}

	// Error case.
	res = reqRes(srv, "POST", "/simple", "some rondom message")

	if code := res.Code; code != 400 {
		t.Errorf("Response code = %v != 400.", code)
	}

	resBody = res.Body.String()
	if resBody != `{"status":"error"}` {
		t.Errorf("Got invalid response: %s.", resBody)
	}

}

func TestSimpleHandlerShow(t *testing.T) {

	d := NewDStore("test", 1000)
	(&Simple{}).SetKind(d)
	defer RefreshDStore(d)

	srv := genSrv(d)

	s := Simple{
		Name: "Ellie",
		Age:  450,
	}

	id, err := s.Save(d)
	if err != nil {
		t.Errorf("Faild to save: %s.", err)
	}

	res := reqRes(srv, "GET", fmt.Sprintf("/simple/%d", id), "")

	if code := res.Code; code != 200 {
		t.Errorf("Response code = %v != 200.", code)
	}

	expected := genJson(&s)
	if resBody := res.Body.String(); !reflect.DeepEqual(resBody, expected) {
		t.Errorf("Expected %v, got %v.", expected, resBody)
	}

	// Error case.

	res = reqRes(srv, "GET", "/simple/1234567", "")
	if code := res.Code; code != 400 {
		t.Errorf("Response code = %v != 400.", code)
	}
}

func TestSimpleHandlerDelete(t *testing.T) {

	d := NewDStore("test", 1000)
	(&Simple{}).SetKind(d)
	defer RefreshDStore(d)

	srv := genSrv(d)

	s := Simple{
		Name: "Akari",
		Age:  15,
	}

	id, err := s.Save(d)
	if err != nil {
		t.Errorf("Faild to save: %s.", err)
	}

	res := reqRes(srv, "DELETE", fmt.Sprintf("/simple/%d", id), "")

	if code := res.Code; code != 200 {
		t.Errorf("Response code = %v != 200.", code)
	}

	if err := (&Simple{}).Find(d, id); err.Error() != "datastore: no such entity" {
		t.Errorf("Expected error not raised: %s.", err)
	}

	// No content case.
	res = reqRes(srv, "DELETE", fmt.Sprintf("/simple/%d", 12234), "")

	if code := res.Code; code != 204 {
		t.Errorf("Response code = %v != 204.", code)
	}
}

func TestSimpleHandlerIndex(t *testing.T) {

	d := NewDStore("test", 1000)
	(&Simple{}).SetKind(d)
	defer RefreshDStore(d)

	srv := genSrv(d)

	createNum := 19

	ss := make([]Simple, createNum)

	for i := 0; i < createNum; i++ {
		ss[i] = Simple{Age: i, Name: "Hinata"}
		if _, err := ss[i].Save(d); err != nil {
			t.Errorf("Faild to save: %s.", err)
		}
	}

	fetchNum := 11

	// The case when limit is specified.
	res := reqRes(srv, "GET", fmt.Sprintf("/simple?limit=%d", fetchNum), "")

	if code := res.Code; code != 200 {
		t.Errorf("Response code = %v != 200.", code)
	}

	js, _ := json.Marshal(ss[0:fetchNum])
	expected := string(js)
	if resBody := res.Body.String(); resBody != expected {
		t.Errorf("Expected %v, got %v.", expected, resBody)
	}

	// The case when limit is NOT specified.
	res = reqRes(srv, "GET", "/simple", "")

	if code := res.Code; code != 200 {
		t.Errorf("Response code = %v != 200.", code)
	}

	js, _ = json.Marshal(ss[0:10])
	expected = string(js)
	if resBody := res.Body.String(); resBody != expected {
		t.Errorf("Expected %v, got %v.", expected, resBody)
	}
}
