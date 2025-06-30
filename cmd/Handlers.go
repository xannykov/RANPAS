package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func MainHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	files := []string{"header.html", "index.html"}

	for _, fname := range files {
		content, err := os.ReadFile(fname)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(content)
	}
}

func InfoHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	files := []string{"header.html", "info.html"}

	for _, fname := range files {
		content, err := os.ReadFile(fname)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(content)
	}
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

	if isForbidden(keyword) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]bool{
			"output-word": true,
		})
		return
	}
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
		r.FormValue("numbers") == "on", 0}

	t := time.Now()

	wordChan := make(chan string)
	go func() {
		wordChan <- fetchRandomWord(w, keyword)
	}()

	crackChan := make(chan float64)
	go func() {
		crackChan <- CrackPassword(&config)
	}()

	randomWordIsGenerated := <-wordChan

	length := config.Length - len([]rune(randomWordIsGenerated))

	password := randomWordIsGenerated + GeneratePassword(length, keyword, &config)
	crackingTime := <-crackChan

	response := map[string]interface{}{
		"password":      password,
		"Cracking Time": crackingTime,
		"Duration":      config.Duration}

	//ХАХАХАХ чзх это ебать
	log.Println("Response: ", response)
	w.Header().Set("Content-Type", "application/json")
	timer1 := time.Since(t)
	log.Println("Evaluated time: ", timer1)
	json.NewEncoder(w).Encode(response)
}
