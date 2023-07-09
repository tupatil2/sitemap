package main

import (
	"flag"

	Sitemap "github.com/tusharr-patil/sitemap/pkg"
)

func main() {
	siteName := flag.String("site", "", "provide the name of the site")
	pages := flag.Int("pages", 10, "pages to crawl")
	flag.Parse()
	Sitemap.GenerateSiteMap(*siteName, *pages)
}
