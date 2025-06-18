package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", MainHandle)
	http.HandleFunc("/static/", StaticHandler)
	http.HandleFunc("/generate", GenerateHandler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
