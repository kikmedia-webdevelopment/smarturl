package utils

import (
	"encoding/base64"
	"math/rand"
)

// This function returns (secure) random bytes
func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func RandomPass(length int64) (string, error) {
	b, err := generateBytes(length)
	return base64.URLEncoding.EncodeToString(b), err
}
