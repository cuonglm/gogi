#gogi - Go client for gitignore.io

[![Build Status](https://travis-ci.org/Gnouc/gogi.svg?branch=master)](https://travis-ci.org/Gnouc/gogi)
[![Go Report Card](https://goreportcard.com/badge/github.com/Gnouc/gogi)](https://goreportcard.com/report/github.com/Gnouc/gogi)

#Why gogi?

Make gitignore client more portable, without relying on the shell, curl, wget or any other http client.

#Installation
```sh
go get -u github.com/Gnouc/gogi
```

#Usage
```sh
import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Gnouc/gogi"
)

func main() {
	gogiClient := gogi.NewHTTPClient(nil)
	resp, _ := gogiClient.List()
	body, err := ioutil.ReadAll(resp.Body)
	if body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
```

#Environment variables

`GOGI_API_URL` to change your gitignore server, default to https://www.gitignore.io

#Author

Cuong Manh Le <cuong.manhle.vn@gmail.com>

#License

See [LICENSE](https://github.com/Gnouc/godt/blob/master/LICENSE)
