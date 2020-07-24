package utils

import (
	"crypto/rand"
	"math/big"
	"unicode"
)

// GenerateRandomString generates a random string with an predefined length
func GenerateRandomString(length int) (string, error) {
	var result string
	for len(result) < length {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		if unicode.IsLetter(rune(n)) {
			result += string(n)
		}
	}
	return result, nil
}
