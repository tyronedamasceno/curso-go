package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := user{"Tyrone", "tyronedamasceno@gmail.com"}
		templates.ExecuteTemplate(w, "index.html", u)
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
