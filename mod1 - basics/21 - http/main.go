package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index page"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
