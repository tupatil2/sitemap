package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	Sitemap "github.com/tusharr-patil/sitemap/pkg"
)

var siteName string

type Response struct {
	XMLResponse string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmpl/main.page.tmpl"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerSubmit(w http.ResponseWriter, r *http.Request) {
	siteName = r.FormValue("siteName")
	log.Println("site name is received:", siteName)
	pages := 10

	xmlResponse := Sitemap.GenerateSiteMap(siteName, pages)

	// Set the response content type as xml text
	w.Header().Set("Content-Type", "text/xml")

	// Write the XML response
	w.Write([]byte(xmlResponse))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/submit", handlerSubmit)
	log.Println("starting server at port 8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = "0.0.0.0:" + port
	}
	http.ListenAndServe(port, nil)
}
