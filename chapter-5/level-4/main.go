package main

func crypt(plaintext, key []byte) []byte {
	result := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		result[i] = plaintext[i] ^ key[i%len(key)]
	}
	return result
}
