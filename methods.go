package gogi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// List list all defined gitignore types
func (c *Client) List() (string, error) {
	path := fmt.Sprintf("%s/list", typePath)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return getResponseBody(resp)
}

// Create create .gitignore content for input type
func (c *Client) Create(typeName string) (string, error) {
	path := fmt.Sprintf("%s/%s", typePath, typeName)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return getResponseBody(resp)
}

func getResponseBody(resp *http.Response) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
