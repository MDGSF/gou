package uhash

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1 calculate sha1 for data
func SHA1(data []byte) (digest string, err error) {
	hasher := sha1.New()
	if _, err := hasher.Write(data); err != nil {
		return "", err
	}
	digest = hex.EncodeToString(hasher.Sum(nil))
	return
}
