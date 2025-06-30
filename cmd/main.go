package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", MainHandle)
	http.HandleFunc("/info", InfoHandle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/generate", GenerateHandler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
