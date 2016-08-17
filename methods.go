package gogi

import (
	"fmt"
	"net/http"
)

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
