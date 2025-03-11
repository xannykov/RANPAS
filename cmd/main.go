package main

import (
	"log"
	"net/http"
)

func MainHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	/* Я этот обработчик въебал по приколу, потому что пока хз как хостить все css файлы не из одного
	мне пока что похуй, потом займусь архитектурой проекта
	*/
	r.URL.Path = r.URL.Path[7:]
	http.FileServer(http.Dir("static")).ServeHTTP(w, r)
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not allowed!", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	keyword := r.FormValue("keyword")

	log.Println("User entered a keyword:", keyword)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", MainHandle)
	http.HandleFunc("/static/", StaticHandler)
	http.HandleFunc("/generate", GenerateHandler)
	log.Fatal(http.ListenAndServe(":3030", nil)) // :8080 не работает, по приколу ему

}
