package uhash

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 calculate md5 for data
func MD5(data []byte) (digest string, err error) {
	hasher := md5.New()
	if _, err := hasher.Write(data); err != nil {
		return "", err
	}
	digest = hex.EncodeToString(hasher.Sum(nil))
	return
}
