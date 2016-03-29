package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"gopkg.in/readline.v1"
)

const (
	xpid    = "pkg00@PASSPORT3496.Rovi"
	custid  = "passport"
	rpr     = "10"
	ectJson = "5"
	baseUrl = "http://roviapi.veveo.net/search"
)

var params map[string]string

func init() {
	params = make(map[string]string)

	params["XPID"] = xpid
	params["custid"] = custid
	params["RPR"] = rpr
	params["ECT"] = ectJson
	//map["W"] = ""
}

func prompt(name, historyFileName string) string {
	prmpt := fmt.Sprintf("%v: ", name)
	rl, err := readline.NewEx(&readline.Config{
		Prompt:      prmpt,
		HistoryFile: "/tmp/" + historyFileName + ".tmp",
	})

	if err != nil {
		panic(err)
	}
	defer rl.Close()

	line, err := rl.Readline()
	if err != nil { // io.EOF
		return "quit"
	}
	return line
}

func searchTerm() (searchString string) {
	fmt.Println("Enter search term (empty string to end):")
	//fmt.Scanln(&searchString)
	searchString = prompt("Search Term", "searchTerm")
	return searchString
}

func getUrl(searchTerm string) string {
	//url := fmt.Sprintf(baseUrl, xpid, custid, rpr, ectJson, searchTerm)
	var Url *url.URL
	Url, err := url.Parse(baseUrl)

	if err != nil {
		panic("Wrong base URL: " + baseUrl)
	}

	parameters := url.Values{}
	for k, v := range params {
		parameters.Add(k, v)
	}

	parameters.Add("W", searchTerm)
	Url.RawQuery = parameters.Encode()

	return Url.String()
}

func main() {

	for {
		searchterm := searchTerm()
		if searchterm == "" {
			break
		}
		url := getUrl(searchterm)
		response, err := http.Get(url)

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			fmt.Println("--------------------begin---------------------------")
			fmt.Println("RQ:" + url)
			fmt.Println("--------------------begin---------------------------")
			fmt.Printf("%s\n", string(contents))
			fmt.Println("---------------------end----------------------------")
		}
	}
}
