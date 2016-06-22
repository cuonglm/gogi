package gogi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	client *Client
	mux    *http.ServeMux
	server *httptest.Server
)

func setUp() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewHTTPClient(nil)
	url, _ := url.Parse(server.URL)
	client.APIURL = url
}

func tearDown() {
	server.Close()
}

func assertEqual(t *testing.T, result interface{}, expect interface{}) {
	if result != expect {
		t.Errorf("Expect (Value: %v) (Type: %T) - Got (Value: %v) (Type: %T)", expect, expect, result, result)
	}
}

func TestNewHTTPClient(t *testing.T) {
	c := NewHTTPClient(nil)

	assertEqual(t, c.UserAgent, ua)
	assertEqual(t, c.APIURL.String(), defaultAPIURL)
}

func TestNewrequest(t *testing.T) {
	c := NewHTTPClient(nil)

	req, _ := c.NewRequest("GET", "/foo", nil)
	assertEqual(t, req.URL.String(), defaultAPIURL+"/foo")
	assertEqual(t, req.Header.Get("User-Agent"), ua)
}

func TestDo(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")

		_, _ = w.Write([]byte("foo"))
	})

	req, _ := client.NewRequest("GET", "/", nil)

	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	res := string(body)
	expected := "foo"
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v - Got %v", expected, res)
	}
}

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
		t.Errorf("Expected %v - Got %v", expected, res)
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
		t.Errorf("Expected %v - Got %v", expected, res)
	}
}
