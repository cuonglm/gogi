package gogi

import (
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

	client, _ = NewHTTPClient()
	url, _ := url.Parse(server.URL)
	client.APIURL = url
}

func tearDown() {
	server.Close()
}

func assertEqual(t *testing.T, result interface{}, expect interface{}) {
	if result != expect {
		t.Fatalf("Expect (Value: %v) (Type: %T) - Got (Value: %v) (Type: %T)", expect, expect, result, result)
	}
}

func TestNewHTTPClient(t *testing.T) {
	c, _ := NewHTTPClient()

	assertEqual(t, c.UserAgent, ua)
	assertEqual(t, c.APIURL.String(), defaultAPIURL)
}

func TestNewrequest(t *testing.T) {
	c, _ := NewHTTPClient()

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
		t.Fatalf("Expected %v - Got %v", expected, res)
	}
}

func TestWithAPIUrl(t *testing.T) {
	c, _ := NewHTTPClient()
	u := "http://cuonglm.xyz"
	if err := WithAPIUrl(u)(c); err != nil {
		t.Fatalf("APIUrl() %+v", err)
	}

	assertEqual(t, c.APIURL.String(), u)
}

func TestWithHTTPClient(t *testing.T) {
	c, _ := NewHTTPClient()

	if err := WithHTTPClient(nil)(c); err == nil {
		t.Fatal("APIUrl() want error, got nil")
	}
}
