package gogi

import (
	"fmt"
	"net/http"
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

	data, err := client.List()
	if err != nil {
		t.Fatalf("List(): %v", err)
	}

	expected := "test list"
	if data != expected {
		t.Fatalf("Expected %v - Got %v", expected, data)
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

	data, err := client.Create("foo")
	if err != nil {
		t.Fatalf("Create(): %v", err)
	}

	expected := "test create foo"
	if data != expected {
		t.Fatalf("Expected %v - Got %v", expected, data)
	}
}
