package data

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomString создаёт случайную строку длины n из латинских букв (A-Z, a-z)
func GenerateRandomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Инициализация генератора случайными числами

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GeneratePhoneNumber генерирует случайный номер в формате +7 (XXX) XXX-XX-XX
func GeneratePhoneNumber() string {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Инициализация генератора случайными числами

	// Генерируем 10 цифр (без +7)
	digits := make([]any, 10)
	for i := range digits {
		digits[i] = rand.Intn(10) // случайная цифра от 0 до 9
	}

	return fmt.Sprintf("+7 (%d%d%d) %d%d%d-%d%d-%d%d", digits...)
}
