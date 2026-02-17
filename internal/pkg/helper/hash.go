package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(code string) string {
	sum := sha256.Sum256([]byte(code))
	return hex.EncodeToString(sum[:])
}