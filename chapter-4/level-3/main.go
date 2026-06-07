package main

import "strings"

func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	result := ""
	for _, c := range text {
		result += getOffsetChar(c, key)
	}
	return result
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	index := strings.IndexRune(alphabet, c)
	if index == -1 {
		return string(c)
	}
	newIndex := ((index + offset) % 26 + 26) % 26
	return string(alphabet[newIndex])
}
