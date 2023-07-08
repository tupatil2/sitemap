package Sitemap

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	Link "github.com/tusharr-patil/html-link-parser"
	"golang.org/x/net/html"
)

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	Comment string   `xml:",comment"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

var site string = ""

var hostName string = ""

// starting point
func Init(siteName string) {

	// site name parsing
	siteLen := len(siteName)
	if siteName[siteLen-1] == '/' {
		site = siteName[0 : siteLen-1]
	} else {
		site = siteName
	}

	// validating the url
	parsedURL, err := url.Parse(site)
	if err != nil {
		log.Fatalln("Error which parsing the siteName URL")
	}
	hostName = parsedURL.Host

	// get the urls
	log.Println("getting the urls")
	urls := getAllUrls(site)

	// generate xml
	log.Println("generating the xml file")
	encodeToXML(urls)
}

// encode the string urls to xml format
func encodeToXML(urls []string) {
	var urlArray []URL

	for _, url := range urls {
		urlArray = append(urlArray, URL{Loc: url})
	}

	urlSet := URLSet{
		XMLNS:   "http://www.sitemaps.org/schemas/sitemap/0.9",
		Comment: "Sitemap generated",
		URLs:    urlArray,
	}

	xmlData, err := xml.MarshalIndent(urlSet, "", "  ")

	if err != nil {
		fmt.Println("Error encoding XML:", err)
	}

	xmlData = append([]byte(xml.Header), xmlData...)

	os.Stdout.Write(xmlData)
}

// gets all the urls related to siteName using bfs
func getAllUrls(siteName string) []string {
	vis := make(map[string]bool)
	q := Queue{}

	var urls []string

	vis[siteName] = true
	q.Enqueue(siteName)
	urls = append(urls, siteName)

	for q.Size() != 0 {
		size := q.Size()
		for size > 0 {
			link := q.Dequeue()
			links := parseLink(link)

			for _, vals := range links {
				childLink := vals.Href
				if childLink == "" {
					continue
				}
				if childLink[0] == '/' {
					parsedURL := ParsePath(childLink)
					if parsedURL == "" || get(vis, parsedURL) {
						continue
					}
					urls = append(urls, parsedURL)
					q.Enqueue(parsedURL)
					vis[parsedURL] = true
				} else {
					parsedURL := ParseSite(childLink)
					if parsedURL == "" || get(vis, parsedURL) {
						continue
					}
					urls = append(urls, parsedURL)
					q.Enqueue(parsedURL)
					vis[parsedURL] = true
				}
			}
			size--
		}
	}

	return urls
}

// parses the string which is URL
func ParseSite(path string) string {
	parsedURL, err := url.Parse(path)
	if err != nil {
		log.Println("Error which parsing the URL")
		return ""
	}

	if parsedURL.Host != hostName || isSectionLink(path) {
		return ""
	}

	return path
}

// parses the strings which starts with "/"
func ParsePath(path string) string {
	if path == "" || path == "/" || isSectionLink(path) {
		return ""
	}

	return site + path
}

// checks if the path is a section link or not
func isSectionLink(path string) bool {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return false
	}
	for _, part := range parts {
		if len(part) != 0 && part[0] == '#' {
			return true
		}
	}
	return false
}

// parses the link and gives all the a tags from the body of that link
func parseLink(link string) []Link.Link {
	log.Println(link)
	response, err := http.Get(link)

	if err != nil {
		fmt.Println("Error wile fetching response from the site")
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Http request failed with status ", response.StatusCode)
		return []Link.Link{}
	}

	doc, err := html.Parse(response.Body)

	if err != nil {
		fmt.Println("error while parsing html", err)
		panic("error while parsing html")
	}

	return Link.GetLinks(doc)
}

// containsKey for map or not
func get(m map[string]bool, key string) bool {
	contains, exists := m[key]
	if !exists {
		return false
	}
	return contains
}
