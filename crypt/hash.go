package crypt

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateRandomSalt() string {
	bytes := make([]byte, 16) // 128 bits
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func HashPasswordWithSalt(password string) (hash, salt string) {
	salt = generateRandomSalt()
	combined := password + salt
	return HashPassword(combined), salt
}

func VerifyPassword(password string, hash string) bool {
	return HashPassword(password) == hash
}

func VerifyPasswordWithSalt(password string, storedHash string, storedSalt string) bool {
	return HashPassword(password+storedSalt) == storedHash
}
