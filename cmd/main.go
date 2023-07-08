package main

import (
	"flag"

	Sitemap "github.com/tusharr-patil/sitemap/pkg"
)

func main() {
	siteName := flag.String("site", "", "a string")
	flag.Parse()
	Sitemap.GenerateSiteMap(*siteName)
}
