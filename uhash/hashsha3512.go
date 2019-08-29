package uhash

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// SHA3512 calculate sha3-512 for data
func SHA3512(data []byte) (digest string, err error) {
	hasher := sha3.New512()
	if _, err := hasher.Write(data); err != nil {
		return "", err
	}
	digest = hex.EncodeToString(hasher.Sum(nil))
	return
}
