package main

import (
	"encoding/hex"
	"math/rand"
)

// GenerateRandomString creates a random hex string of length n.
func GenerateRandomString(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
