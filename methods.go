package gogi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ListJsonItem struct {
	Contents string `json:"contents"`
	FileName string `json:"fileName"`
	Key      string `json:"key"`
	Name     string `json:"name"`
}

// List lists all defined gitignore types with lines format
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

// ListJson lists all defined gitignore types with json format.
func (c *Client) ListJson() (map[string]*ListJsonItem, error) {
	path := fmt.Sprintf("%s/list", typePath)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("format", "json")
	req.URL.RawQuery = q.Encode()

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res map[string]*ListJsonItem
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// Create creates .gitignore file content for given input type.
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
