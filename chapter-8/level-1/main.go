package main

func feistel(msg []byte, roundKeys [][]byte) []byte {
	// Split message into left and right halves
	mid := len(msg) / 2
	left := msg[:mid]
	right := msg[mid:]

	// Perform Feistel rounds
	for _, roundKey := range roundKeys {
		// Hash the right half with the round key
		hashResult := hash(right, roundKey, len(left))
		// XOR left with hash result to get new right
		newRight := xor(left, hashResult)
		// Right becomes new left
		left = right
		right = newRight
	}

	// Swap and combine left and right halves
	return append(right, left...)
}
