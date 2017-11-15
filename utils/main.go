package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// RandomHex return a random hex string
func RandomHex(n int) string {
	return hex.EncodeToString(RandomBytes(n))
}

// RandomBytes return a random bytes
func RandomBytes(n int) []byte {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return bytes
}

// Hash return a sha256 hex string
func Hash(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}
