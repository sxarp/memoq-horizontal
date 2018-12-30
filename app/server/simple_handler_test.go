package main

import (
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

	res := reqRes(srv, "POST", "/simple/create", body)

	if code := res.Code; code != 200 {
		t.Errorf("response code = %v != 200", code)
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
	res = reqRes(srv, "POST", "/simple/create", "some rondom message")

	if code := res.Code; code != 400 {
		t.Errorf("response code = %v != 400", code)
	}

	resBody = res.Body.String()
	if resBody != `{"status":"error"}` {
		t.Errorf("Got invalid response: %s.", resBody)
	}

}
