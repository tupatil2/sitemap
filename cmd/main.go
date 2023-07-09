package main

import (
	"flag"

	Sitemap "github.com/tusharr-patil/sitemap/pkg"
)

func main() {
	siteName := flag.String("site", "", "provide the name of the site")
	flag.Parse()
	Sitemap.GenerateSiteMap(*siteName)
}
