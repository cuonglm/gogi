package gogi

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	version       = "0.0.2"
	ua            = "gogi/" + version
	defaultAPIURL = "https://www.gitignore.io"
	typePath      = "/api"
)

// Client for querying API
type Client struct {
	client    *http.Client
	UserAgent string
	APIURL    *url.URL
}

// NewHTTPClient create new gogi client
func NewHTTPClient(options ...func(*Client) error) (*Client, error) {

	c := &Client{
		client:    http.DefaultClient,
		UserAgent: ua,
	}

	// Default API url
	err := APIUrl(defaultAPIURL)(c)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// APIUrl sets the API url option for gogi client
func APIUrl(u string) func(*Client) error {
	return func(c *Client) error {
		apiURL, err := url.Parse(u)
		if err != nil {
			return err
		}

		c.APIURL = apiURL

		return nil
	}
}

// HTTPClient sets the client option for gogi client
func HTTPClient(client *http.Client) func(*Client) error {
	return func(c *Client) error {
		if client == nil {
			return errors.New("client is nil")
		}

		c.client = client

		return nil
	}
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
