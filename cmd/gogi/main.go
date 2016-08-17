package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Gnouc/gogi"
)

var (
	listFlag   bool
	createFlag string
	searchFlag string
	gogiClient *gogi.Client
)

func init() {
	flag.BoolVar(&listFlag, "list", false, "List all defined types")
	flag.StringVar(&createFlag, "create", "", "Create .gitignore content for given types")
	flag.StringVar(&searchFlag, "search", "", "Show all types match string")

	gogiClient = gogi.NewHTTPClient(nil)
}

func main() {
	flag.Parse()

	n := flag.NFlag()
	switch {
	case n == 1:
		switch {
		case listFlag:
			list()
		case createFlag != "":
			create(createFlag)
		case searchFlag != "":
			search(searchFlag)
		}
	case n >= 2:
		fmt.Println("Only one action allow.")
		fmt.Println()
		fallthrough
	default:
		flag.Usage()
	}
}

func list() {
	resp, _ := gogiClient.List()
	data := extractResponse(resp)
	fmt.Println(data)
}

func create(s string) {
	resp, _ := gogiClient.Create(s)
	data := extractResponse(resp)
	fmt.Println(data)
}

func search(s string) {
	resp, _ := gogiClient.List()
	data := extractResponse(resp)
	data = strings.Replace(data, "\n", ",", -1)

	for _, v := range strings.Split(data, ",") {
		if strings.Contains(v, s) {
			fmt.Println(v)
		}
	}
}

func extractResponse(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
