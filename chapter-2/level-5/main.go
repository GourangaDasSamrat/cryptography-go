package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	hexParts := make([]string, len(b))
	for i, byte := range b {
		hexParts[i] = fmt.Sprintf("%02x", byte)
	}
	return strings.Join(hexParts, ":")
}

func getBinaryString(b []byte) string {
	binParts := make([]string, len(b))
	for i, byte := range b {
		binParts[i] = fmt.Sprintf("%08b", byte)
	}
	return strings.Join(binParts, ":")
}
