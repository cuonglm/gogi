package gogi

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestNewHTTPClient(t *testing.T) {
	c, _ := NewHTTPClient()

	assert.Equal(t, ua, c.UserAgent)
	assert.Equal(t, defaultAPIURL, c.APIURL.String())
}

func TestNewrequest(t *testing.T) {
	c, _ := NewHTTPClient()
	req, _ := c.NewRequest("GET", "/foo", nil)

	assert.Equal(t, defaultAPIURL+"/foo", req.URL.String())
	assert.Equal(t, ua, req.Header.Get("User-Agent"))
}

func TestDo(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		_, _ = w.Write([]byte("foo"))
	})

	req, err := client.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	res := string(body)
	expected := "foo"
	assert.Equal(t, expected, res)
}

func TestWithAPIUrl(t *testing.T) {
	c, _ := NewHTTPClient()
	u := "https://cuonglm.xyz"
	require.NoError(t, WithAPIUrl(u)(c))
	assert.Equal(t, u, c.APIURL.String())
}

func TestWithHTTPClient(t *testing.T) {
	c, _ := NewHTTPClient()
	require.Error(t, WithHTTPClient(nil)(c))
}
