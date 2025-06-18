package main

import (
	"log"
	"math/big"
	"net/http"
	"strconv"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      = "0123456789"
	symbols      = "!@#$%^&*()_+-=[]{}|;:,.<>?/~"
)

func GeneratePassword(length int, Uselower, Useupper, Usenumber, Usesymbol bool) string {

	// добавить проверку на длительность взлома(формула андерсона) в отдельную функ-ию ебать которую

	chars := ""

	if Uselower {
		chars += lowerLetters
	}

	if Useupper {
		chars += upperLetters
	}

	if Usenumber {
		chars += numbers
	}

	if Usesymbol {
		chars += symbols
	}
	if chars == "" {
		return chars
	}

	password := make([]byte, length)
	//Обработать если длинна стала меньше нуля
	for i := range password {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		password[i] = chars[randIndex.Int64()]
	}

	return string(password)
}

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
	log.Println("Entered keyword: ", keyword)
	lenght, _ := strconv.Atoi(r.FormValue("passwordLength"))

	if lenght < 7 {
		lenght = 7
	}
	if lenght > 20 {
		lenght = 20
	}

	lenght -= len(keyword)

	useLower := r.FormValue("smallLetters") == "on"
	useUpper := r.FormValue("bigLetters") == "on"
	useSymbols := r.FormValue("symbols") == "on"
	useNumbers := r.FormValue("numbers") == "on"

	password := keyword + GeneratePassword(lenght, useLower, useUpper, useNumbers, useSymbols)

	response := map[string]string{"password": password}

	//ХАХАХАХ чзх это ебать
	log.Println("Response: ", response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", MainHandle)
	http.HandleFunc("/static/", StaticHandler)
	http.HandleFunc("/generate", GenerateHandler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
