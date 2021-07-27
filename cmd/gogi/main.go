package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cuonglm/gogi"
)

var (
	listFlag   bool
	jsonFlag   bool
	createFlag string
	searchFlag string
	gogiClient *gogi.Client
	err        error
	apiURL     string
)

func init() {
	flag.BoolVar(&listFlag, "list", false, "List all defined types")
	flag.BoolVar(&jsonFlag, "json", false, "List with JSON format, use with -list")
	flag.StringVar(&createFlag, "create", "", "Create .gitignore content for given types")
	flag.StringVar(&searchFlag, "search", "", "Show all types match string")

	apiURL = os.Getenv("GOGI_API_URL")
	if apiURL == "" {
		gogiClient, _ = gogi.NewHTTPClient()
		return
	}
	gogiClient, err = gogi.NewHTTPClient(gogi.WithAPIUrl(apiURL))
	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	switch flag.NFlag() {
	case 1:
		// ok
	case 2:
		if listFlag && jsonFlag {
			break // ok
		}
		fallthrough
	default:
		printError("Only one action allow.")
		flag.Usage()
		os.Exit(1)
	}
	switch {
	case listFlag && jsonFlag:
		listJson()
	case jsonFlag:
		printError("-json flag must be used with -list")
		os.Exit(1)
	case listFlag:
		list()
	case createFlag != "":
		create(createFlag)
	case searchFlag != "":
		search(searchFlag)
	}
}

func list() {
	data, err := gogiClient.List()
	if err != nil {
		log.Fatal(err)
	}

	_, _ = fmt.Fprintln(os.Stderr, data)
}

func listJson() {
	data, err := gogiClient.ListJson()
	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(data)
}

func create(s string) {
	data, err := gogiClient.Create(s)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = fmt.Fprintln(os.Stderr, data)
}

func search(s string) {
	data, err := gogiClient.List()
	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}
	data = strings.Replace(data, "\n", ",", -1)

	for _, v := range strings.Split(data, ",") {
		if strings.Contains(v, s) {
			fmt.Println(v)
		}
	}
}

func printError(msg string) {
	_, _ = fmt.Fprintln(os.Stderr, msg)
}
