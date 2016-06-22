package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Gnouc/gogi"
)

var (
	listFlag   bool
	createFlag string
	gogiClient *gogi.Client
)

func init() {
	flag.BoolVar(&listFlag, "list", false, "List all defined types")
	flag.StringVar(&createFlag, "create", "", "Create .gitignore file for given types")

	gogiClient = gogi.NewHTTPClient(nil)
}

func main() {
	flag.Parse()

	switch flag.NFlag() {
	case 1:
		if listFlag {
			list()
		}
		if createFlag != "" {
			create(createFlag)
		}
	case 2:
		log.Println("Only one action allow.")
		fallthrough
	default:
		flag.Usage()
	}
}

func list() {
	resp, _ := gogiClient.List()
	printResponse(resp)
}

func create(s string) {
	resp, _ := gogiClient.Create(s)
	printResponse(resp)
}

func printResponse(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	os.Exit(0)
}
