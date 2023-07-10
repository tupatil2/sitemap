package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	Sitemap "github.com/tusharr-patil/sitemap/pkg"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmpl/main.page.tmpl"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	siteName := flag.String("site", "", "provide the name of the site")
	pages := flag.Int("pages", 10, "pages to crawl")
	flag.Parse()
	Sitemap.GenerateSiteMap(*siteName, *pages)

	log.Println("starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}
