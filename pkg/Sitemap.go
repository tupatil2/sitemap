package Sitemap

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

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
	Loc      string  `xml:"loc"`
	Lastmod  string  `xml:"lastmod,omitempty"`
	Priority float64 `xml:"priority,omitempty"`
}

var site string = ""
var hostName string = ""

// starting point
func GenerateSiteMap(siteName string, pages int64, priority float64, modifieddate bool) string {

	// site name parsing
	siteLen := len(siteName)
	if siteLen == 0 {
		log.Println("no sitename provided")
		return ""
	}

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
	var lastmod string = ""
	if modifieddate {
		now := time.Now().UTC()
		lastmod = string(now.Format("2006-01-02T15:04:05Z07:00"))
	}

	// get the urls
	log.Println("getting the urls")
	urls := getAllUrls(site, pages)

	// generate xml
	return encodeToXML(urls, priority, lastmod)
}

// encode the string urls to xml format
func encodeToXML(urls []string, priority float64, lastmod string) string {
	var urlArray []URL

	for _, url := range urls {
		url := URL{Loc: url}
		if lastmod != "" {
			url.Lastmod = lastmod
		}
		if priority != -1.0 {
			url.Priority = priority
		}
		urlArray = append(urlArray, url)
	}

	urlSet := URLSet{
		XMLNS:   "http://www.sitemaps.org/schemas/sitemap/0.9",
		Comment: "Sitemap generated",
		URLs:    urlArray,
	}

	xmlData, err := xml.MarshalIndent(urlSet, "", "  ")

	if err != nil {
		log.Println("Error encoding XML:", err)
	}

	xmlData = append([]byte(xml.Header), xmlData...)

	log.Println("generated the xml file")

	return string(xmlData)
}

// gets all the urls related to siteName using bfs
func getAllUrls(siteName string, pages int64) []string {
	vis := make(map[string]bool)
	q := Queue{}

	var urls []string

	vis[siteName] = true
	q.enqueue(siteName)
	urls = append(urls, siteName)
	depth := pages

	for q.size() != 0 {
		size := q.size()
		for size > 0 {
			link := q.dequeue()
			links := parseLink(link)

			for _, vals := range links {
				childLink := vals.Href
				if childLink == "" {
					continue
				}
				if childLink[0] == '/' {
					parsedURL := parsePath(childLink)
					if parsedURL == "" || get(vis, parsedURL) {
						continue
					}
					urls = append(urls, parsedURL)
					q.enqueue(parsedURL)
					vis[parsedURL] = true
				} else {
					parsedURL := parseSite(childLink)
					if parsedURL == "" || get(vis, parsedURL) {
						continue
					}
					urls = append(urls, parsedURL)
					q.enqueue(parsedURL)
					vis[parsedURL] = true
				}
			}
			size--
		}
		depth--
		if depth == 0 {
			break
		}
	}

	return urls
}

// parses the string which is URL
func parseSite(path string) string {
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
func parsePath(path string) string {
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
		log.Println("Error wile fetching response from the site")
		panic("Error wile fetching response from the site")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("Http request failed with status ", response.StatusCode)
		panic("Http request failed")
	}

	doc, err := html.Parse(response.Body)

	if err != nil {
		log.Println("error while parsing html", err)
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
