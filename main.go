package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	Sitemap "github.com/tusharr-patil/sitemap/pkg"
)

var siteName string
var priority float64 = -1.0
var modifieddate bool = false

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
	if r.FormValue("modifieddate") != "none" {
		modifieddate = true
	}
	if r.FormValue("priority") != "none" {
		priority, _ = strconv.ParseFloat(r.FormValue("priority"), 64)
	} else {
		priority = -1.0
	}

	pages, _ := strconv.ParseInt(r.FormValue("pages"), 10, 64)

	log.Println("site name:", siteName)
	log.Println("modified date: ", r.FormValue("modifieddate"))
	log.Println("priority: ", r.FormValue("priority"))
	log.Println("pages: ", r.FormValue("pages"))

	xmlResponse := Sitemap.GenerateSiteMap(siteName, pages, priority, modifieddate)

	// Set the response content type as xml text
	w.Header().Set("Content-Type", "text/xml")

	// Write the XML response
	w.Write([]byte(xmlResponse))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/submit", handlerSubmit)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = "0.0.0.0:" + port
	}
	log.Printf("starting server at port %v \n", port)
	http.ListenAndServe(port, nil)
}
