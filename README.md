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

## As library
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

## As binary:
```sh
$ go get github.com/Gnouc/gogi/gogi
$ gogi
Usage of gogi:
  -create string
    	Create .gitignore content for given types
  -list
    	List all defined types
  -search string
    	Show all types match string
```

## Using docker

### Using `gnouc/gogi` image
```sh
$ docker pull gnouc/gogi
$ docker run --rm gnouc/gogi -search python
ipythonnotebook
python
```

### Building your own image

Building builder image
```sh
docker build -t gogi-builder -f Dockerfile.build .
```

Building binary image
```sh
docker run --rm gogi-builder | docker build -t gogi -f Dockerfile.run -
```

#Environment variables

`GOGI_API_URL` to change your gitignore server, default to https://www.gitignore.io

#Author

Cuong Manh Le <cuong.manhle.vn@gmail.com>

#License

See [LICENSE](https://github.com/Gnouc/godt/blob/master/LICENSE)
