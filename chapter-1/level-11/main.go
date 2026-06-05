package main

import (
	"encoding/hex"
	"math/rand"
)

func generateRandomKey(length int) (string, error) {
	randReader := rand.New(rand.NewSource(0))

	// Generate random bytes
	bytes := make([]byte, length)
	randReader.Read(bytes)

	// Return hex-encoded string
	return hex.EncodeToString(bytes), nil
}
