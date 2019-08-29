package utils

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

// GetRandomBytes generate random byte slice
// @param n[in]: the length of random string
func GetRandomBytes(n int, alphabets ...byte) []byte {
	if n <= 0 {
		return []byte{}
	}

	if len(alphabets) == 0 {
		alphabets = alphaNum
	}

	var bytes = make([]byte, n)

	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}

	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}

// GetRandomString generate random string
// @param n[in]: the length of random string
func GetRandomString(n int, alphabets ...byte) string {
	return string(GetRandomBytes(n, alphabets...))
}
