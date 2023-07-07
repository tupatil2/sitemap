package main

import (
	"fmt"
	"net/http"

	Link "github.com/tusharr-patil/html-link-parser"
	"golang.org/x/net/html"
)

func main() {
	siteName := "https://www.calhoun.io/"

	response, err := http.Get(siteName)

	if err != nil {
		fmt.Println("Error wile fetching response from the site")
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Http request failed with status ", response.StatusCode)
		return
	}

	doc, err := html.Parse(response.Body)

	if err != nil {
		fmt.Println("error while parsing html", err)
		return
	}

	link := Link.GetLinks(doc)
	fmt.Println(link)
}
