package main

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	result := masterKey
	for i := 0; i < 4; i++ {
		result[i] ^= byte(roundNumber)
	}
	return result
}
