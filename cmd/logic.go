package main

import (
	"crypto/rand"
	"log"
	"math"
	"math/big"
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

func GetRandomWords(words []string) string {
	if len(words) == 0 {
		return ""
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
	if err != nil {
		log.Println("Error generating random index:", err)
		return words[0] // В случае ошибки возвращаем первый элемент
	}
	return words[n.Int64()]
}
