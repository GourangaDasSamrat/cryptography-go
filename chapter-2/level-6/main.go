package main

import (
	"encoding/hex"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
	parts := strings.Split(s, ":")
	result := make([]byte, len(parts))

	for i, part := range parts {
		b, err := hex.DecodeString(part)
		if err != nil || len(b) != 1 {
			return nil, err
		}
		result[i] = b[0]
	}

	return result, nil
}
