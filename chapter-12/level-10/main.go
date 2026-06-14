package main

import "math/bits"

func hash(input []byte) [4]byte {
	// Step 1: Rotate bits left 3 bits
	rotated := make([]byte, len(input))
	for i, b := range input {
		rotated[i] = bits.RotateLeft8(b, 3)
	}

	// Step 2: Shift bits left 2 bits
	shifted := make([]byte, len(rotated))
	for i, b := range rotated {
		shifted[i] = b << 2
	}

	// Step 3: Create final array
	final := [4]byte{}

	// Step 4: XOR each byte with corresponding byte in final array
	for i, b := range shifted {
		final[i%4] ^= b
	}

	return final
}
