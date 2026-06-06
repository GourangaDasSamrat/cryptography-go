package main

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	key := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i++ {
		key[i] = encrypted[i] ^ decrypted[i]
	}
	return key, nil
}
