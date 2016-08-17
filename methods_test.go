package gogi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	setUp()
	defer tearDown()

	path := fmt.Sprintf("%s/list", typePath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")

		_, _ = w.Write([]byte("test list"))
	})

	resp, err := client.List()
	if err != nil {
		t.Fatalf("List(): %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("List(): %v", err)
	}

	res := string(body)
	expected := "test list"
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Expected %v - Got %v", expected, res)
	}
}

func TestCreate(t *testing.T) {
	setUp()
	defer tearDown()

	path := fmt.Sprintf("%s/foo", typePath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")

		_, _ = w.Write([]byte("test create foo"))
	})

	resp, err := client.Create("foo")
	if err != nil {
		t.Fatalf("Create(): %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Create(): %v", err)
	}

	res := string(body)
	expected := "test create foo"
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Expected %v - Got %v", expected, res)
	}
}
