package main

import "fmt"

func getBlockSize(keyLen, cipherType int) (int, error) {
	switch cipherType {
	case typeAES:
		// AES supports 128, 192, 256 bit keys (16, 24, 32 bytes)
		if keyLen == 16 || keyLen == 24 || keyLen == 32 {
			return 16, nil
		}
		return 0, fmt.Errorf("invalid AES key length: %d", keyLen)
	case typeDES:
		// DES supports 64 bit key (8 bytes)
		if keyLen == 8 {
			return 8, nil
		}
		return 0, fmt.Errorf("invalid DES key length: %d", keyLen)
	default:
		return 0, fmt.Errorf("invalid cipher type: %d", cipherType)
	}
}
