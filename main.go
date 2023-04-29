package main

import (
	"html/template"
	"log"
	"net/http"
)

type IndexPageData struct {
	PageTitle string
}

type SubPageData struct {
	PageTitle string
}

func main() {
	index_tmpl := template.Must(template.ParseFiles("static/index.html"))
	subpage_tmpl := template.Must(template.ParseFiles("static/subpage.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := IndexPageData{
			PageTitle: "This is main page!",
		}
		index_tmpl.Execute(w, data)
	})

	http.HandleFunc("/subpage/", func(w http.ResponseWriter, r *http.Request) {
		data := SubPageData{
			PageTitle: "This is subpage!",
		}
		subpage_tmpl.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
