package gogi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	setUp()
	defer tearDown()

	path := fmt.Sprintf("%s/list", typePath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		_, _ = w.Write([]byte("test list"))
	})

	data, err := client.List()
	require.NoError(t, err)

	expected := "test list"
	assert.Equal(t, expected, data)
}

func TestListJson(t *testing.T) {
	setUp()
	defer tearDown()

	goItem := &ListJsonItem{
		Contents: "\n### Go ###\n# Binaries for programs and plugins\n*.exe\n*.exe~\n*.dll\n*.so\n*.dylib\n\n# Test binary, built with `go test -c`\n*.test\n\n# Output of the go coverage tool, specifically when used with LiteIDE\n*.out\n\n# Dependency directories (remove the comment below to include it)\n# vendor/\n\n### Go Patch ###\n/vendor/\n/Godeps/\n",
		FileName: "Go.gitignore",
		Key:      "go",
		Name:     "go",
	}
	path := fmt.Sprintf("%s/list", typePath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "json", r.URL.Query().Get("format"))

		res := map[string]*ListJsonItem{
			"go": goItem,
		}
		assert.NoError(t, json.NewEncoder(w).Encode(res))
	})

	data, err := client.ListJson()
	require.NoError(t, err)

	assert.Len(t, data, 1)
	assert.Equal(t, goItem, data["go"])
}

func TestCreate(t *testing.T) {
	setUp()
	defer tearDown()

	path := fmt.Sprintf("%s/foo", typePath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		_, _ = w.Write([]byte("test create foo"))
	})

	data, err := client.Create("foo")
	require.NoError(t, err)

	expected := "test create foo"
	assert.Equal(t, expected, data)
}
