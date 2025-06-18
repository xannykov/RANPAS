package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func MainHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = r.URL.Path[7:]
	http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	//w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
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
	lenghtAllPassword, _ := strconv.Atoi(r.FormValue("passwordLength"))

	if lenghtAllPassword < MinPasswordLength {
		lenghtAllPassword = MinPasswordLength
	}
	if lenghtAllPassword > MaxPasswordLength {
		lenghtAllPassword = MaxPasswordLength
	}

	config := PasswordConfig{
		lenghtAllPassword, r.FormValue("lowLetters") == "on",
		r.FormValue("bigLetters") == "on",
		r.FormValue("symbols") == "on",
		r.FormValue("numbers") == "on", Hour}

	encodedKeyword := url.QueryEscape(keyword)

	urlAddress := fmt.Sprintf("http://localhost:8000/associations?word=%s", encodedKeyword)

	respFromFastAPI, err := http.Get(urlAddress)
	if err != nil {
		http.Error(w, "StatusInternalServerError", http.StatusInternalServerError)
	}
	defer respFromFastAPI.Body.Close()

	body, _ := io.ReadAll(respFromFastAPI.Body)

	log.Println("Response from FastAPI:", string(body))

	words := make([]string, len(body))
	err = json.Unmarshal(body, &words)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		log.Println("Error parsing JSON:", err)
		return
	}
	log.Println(words)

	randomWord := GetRandomWords(words)
	length := config.Length - len([]rune(randomWord))

	password := randomWord + GeneratePassword(length, keyword, &config)
	crackingTime := CrackPassword(&config)

	response := map[string]interface{}{
		"password":      password,
		"Cracking Time": crackingTime,
		"Duration":      config.Duration}

	//ХАХАХАХ чзх это ебать
	log.Println("Response: ", response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
