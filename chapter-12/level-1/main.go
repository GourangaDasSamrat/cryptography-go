package main

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func newHasher() *hasher {
	return &hasher{
		hash: sha256.New(),
	}
}

func (h *hasher) Write(password string) {
	h.hash.Write([]byte(password))
}

func (h *hasher) GetHex() string {
	return hex.EncodeToString(h.hash.Sum(nil))
}
