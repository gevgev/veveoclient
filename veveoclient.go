package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/readline.v1"
)

const (
	xpid    = "pkg00@PASSPORT3496.Rovi"
	custid  = "passport"
	rpr     = 10
	ectJson = 5
	baseUrl = "http://roviapi.veveo.net/search?XPID=%s&custid=%s&RPR=%d&ECT=%d&W=%s"
)

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
	url := fmt.Sprintf(baseUrl, xpid, custid, rpr, ectJson, searchTerm)
	return url
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
			fmt.Printf("%s\n", string(contents))
			fmt.Println("---------------------end----------------------------")
		}
	}
}
