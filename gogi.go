package gogi

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const (
	version       = "0.0.1"
	ua            = "gogi/" + version
	defaultAPIURL = "https://www.gitignore.io"
	typePath      = "/api"
	envAPIURL     = "GOGI_API_URL"
)

// Client for querying API
type Client struct {
	client    *http.Client
	UserAgent string
	APIURL    *url.URL
}

// NewHTTPClient create new gogi client
func NewHTTPClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	apiURLStr := os.Getenv(envAPIURL)
	if apiURLStr == "" {
		apiURLStr = defaultAPIURL
	}
	apiURL, err := url.Parse(apiURLStr)
	if err != nil {
		panic(err)
	}

	c := &Client{
		client:    client,
		UserAgent: ua,
		APIURL:    apiURL,
	}

	return c
}

// NewRequest create new http request
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	relPath, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.APIURL.ResolveReference(relPath)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Do make an http request
func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	resp, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// List list all defined gitignore types
func (c *Client) List() (*http.Response, error) {
	path := fmt.Sprintf("%s/list", typePath)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Create create .gitignore content for input type
func (c *Client) Create(typeName string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", typePath, typeName)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
