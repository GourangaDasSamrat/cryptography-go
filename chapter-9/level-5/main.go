package main

import "math/big"

// nonceStrength returns the number of possible nonce values (2^bits)
func nonceStrength(nonce []byte) int {
	bits := len(nonce) * 8
	strength := big.NewInt(1)
	strength.Lsh(strength, uint(bits))
	return int(strength.Int64())
}
