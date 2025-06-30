package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	lowerLetters      = "abcdefghijklmnopqrstuvwxyz"
	upperLetters      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers           = "0123456789"
	symbols           = "!@#$%^&*()_+-=[]{}|;:,.<>?/~"
	MinPasswordLength = 7
	MaxPasswordLength = 20
)
const (
	Hour = 1 + iota
	Day
	Week
	Mouth
	Year
	Uncountable
)

type PasswordConfig struct {
	Length     int
	UseLower   bool
	UseUpper   bool
	UseNumbers bool
	UseSymbols bool
	Duration   int
}

func GeneratePassword(length int, keyword string, config *PasswordConfig) string {

	chars := ""

	if config.UseLower {
		chars += lowerLetters
	}

	if config.UseUpper {
		chars += upperLetters
	}

	if config.UseNumbers {
		chars += numbers
	}

	if config.UseSymbols {
		chars += symbols
	}
	if chars == "" {
		return chars
	}

	if length <= 0 {
		return keyword
	}

	passwordChars := make([]byte, length)
	for i := range passwordChars {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		passwordChars[i] = chars[randIndex.Int64()]
	}

	return string(passwordChars)
}

func CrackPassword(config *PasswordConfig) float64 {
	Power := 0
	if config.UseLower {
		Power += len(lowerLetters)
	}
	if config.UseUpper {
		Power += len(upperLetters)
	}

	if config.UseNumbers {
		Power += len(numbers)
	}

	if config.UseSymbols {
		Power += len(symbols)
	}

	V := math.Pow(10, 10)
	time := (math.Pow(float64(Power), float64(config.Length))) / V
	time /= 3600

	if time > 24 {
		time /= 24
		config.Duration = Day
	}

	if time > 7 {
		time /= 7
		config.Duration = Week
	}
	if time > 31 {
		time /= 31
		config.Duration = Mouth
	}
	if time > 12 {
		time /= 12
		config.Duration = Year
	}
	if time > 999 {
		time = 999
		config.Duration = Uncountable
	}
	return time
}

func isForbidden(keyword string) bool {
	forbidden := make([]string, 0, 141)
	fileName := "forbidden_words.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return true
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		forbidden = append(forbidden, scanner.Text())
	}
	if err2 := scanner.Err(); err2 != nil {
		log.Fatal(err2)
		return true
	}

	keyword = strings.ToLower(keyword)

	for _, forb := range forbidden {
		if keyword == forb {
			return true
		}
	}
	return false
}

func fetchRandomWord(w http.ResponseWriter, keyword string) string {
	encodedKeyword := url.QueryEscape(keyword)

	urlAddress := fmt.Sprintf("http://localhost:8000/associations?word=%s", encodedKeyword)

	respFromFastAPI, err := http.Get(urlAddress)
	if err != nil {
		http.Error(w, "StatusInternalServerError", http.StatusInternalServerError)
	}
	defer respFromFastAPI.Body.Close()

	body, _ := io.ReadAll(respFromFastAPI.Body)

	log.Println("Response from FastAPI:", string(body))

	var words []string
	err = json.Unmarshal(body, &words)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		log.Println("Error parsing JSON:", err)
		return ""
	}
	return words[0]
}
