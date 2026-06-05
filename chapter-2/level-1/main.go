package main

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	if bits > 7 {
		return ""
	}
	return string(base8Alphabet[bits])
}
