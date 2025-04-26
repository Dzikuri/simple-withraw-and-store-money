package util

import (
	"crypto/rand"
	"math/big"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Get the salt cost factor from the environment variable
	costFactor := os.Getenv("BCRYPT_SALT")
	if costFactor == "" {
		costFactor = "10" // Default to cost factor of 10 if not provided
	}

	// Convert the cost factor to int
	cost, err := strconv.Atoi(costFactor)
	if err != nil {
		return "", err
	}

	// Generate a salted hash for the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}

func IsValidPassword(password string) bool {
	// Periksa panjang password minimal 8 karakter
	if len(password) < 8 {
		return false
	}

	// Regular expression untuk huruf kapital dan simbol
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasSymbol := regexp.MustCompile(`[!@#~$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+`).MatchString(password)

	// Return true jika ada huruf kapital dan simbol
	return hasUppercase && hasSymbol
}

func GenerateRandomPassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
	password := make([]byte, length)
	for i := range password {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[randomIndex.Int64()]
	}
	return string(password)
}
