package main

import "fmt"

func sBox(b byte) (byte, error) {
	if b > 15 {
		return 0, fmt.Errorf("input must be 4-bit value (0-15), got %d", b)
	}

	// S-box lookup table: 4-bit input -> 2-bit output
	sbox := [16]byte{0, 2, 1, 3, 2, 1, 3, 0, 1, 3, 2, 0, 1, 2, 3, 0}

	return sbox[b], nil
}
