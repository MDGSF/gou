package uhash

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 calculate sha256 for data
func SHA256(data []byte) (digest string, err error) {
	hasher := sha256.New()
	if _, err := hasher.Write(data); err != nil {
		return "", err
	}
	digest = hex.EncodeToString(hasher.Sum(nil))
	return
}
