package main

func padWithZeros(block []byte, desiredSize int) []byte {
	padded := make([]byte, desiredSize)
	copy(padded, block)
	return padded
}
