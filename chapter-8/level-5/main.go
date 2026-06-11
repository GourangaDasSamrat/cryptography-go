package main

import (
	"crypto/cipher"
	"crypto/des"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedPlaintext := padMsg(plaintext, block.BlockSize())

	iv := make([]byte, block.BlockSize())

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	return append(iv, ciphertext...), nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	paddedLen := blockSize * ((len(plaintext) + blockSize - 1) / blockSize)
	return padWithZeros(plaintext, paddedLen)
}
