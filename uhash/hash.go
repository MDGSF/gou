// MIT License
//
// Copyright (c) 2019 Huang Jian
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package uhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io/ioutil"
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

// GetFileDigest Returns the hex digest of a file.
func GetFileDigest(filename, algo string) (digest string, err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return GetDataDigest(content, algo)
}
