package uhash

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512 calculate sha512 for data
func SHA512(data []byte) (digest string, err error) {
	hasher := sha512.New()
	if _, err := hasher.Write(data); err != nil {
		return "", err
	}
	digest = hex.EncodeToString(hasher.Sum(nil))
	return
}
