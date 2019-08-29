package uhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"strings"

	"golang.org/x/crypto/sha3"
)

/*
GetDataDigest Returns the hex digest of some data.
@param data: calculate hash for data.
@param algo: md5, sha1, sha256, sha512
*/
func GetDataDigest(data []byte, algo string) (digest string, err error) {
	var hasher hash.Hash

	switch strings.ToLower(algo) {
	case "md5":
		hasher = md5.New()
	case "sha1":
		hasher = sha1.New()
	case "sha256":
		hasher = sha256.New()
	case "sha512":
		hasher = sha512.New()
	case "sha3-512":
		hasher = sha3.New512()
	default:
		err = errors.New("invalid hash algorithm")
		return
	}

	hasher.Write(data)
	digest = hex.EncodeToString(hasher.Sum(nil))
	return
}
