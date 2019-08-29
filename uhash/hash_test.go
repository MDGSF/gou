package uhash

import (
	"log"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDataDigestMd5(t *testing.T) {
	digest, err := GetDataDigest([]byte("huangjian"), "md5")
	assert.Equal(t, nil, err, "they should be equal")

	cmd := exec.Command("/bin/sh", "-c", `echo -n huangjian|openssl md5`)
	output, err := cmd.Output()
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, true, strings.Contains(string(output), digest), "they should be equal")
}

func TestGetDataDigestSha1(t *testing.T) {
	digest, err := GetDataDigest([]byte("huangjian"), "sha1")
	assert.Equal(t, nil, err, "they should be equal")

	cmd := exec.Command("/bin/sh", "-c", `echo -n huangjian|openssl sha1`)
	output, err := cmd.Output()
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, true, strings.Contains(string(output), digest), "they should be equal")
}

func TestGetDataDigestSha256(t *testing.T) {
	digest, err := GetDataDigest([]byte("huangjian"), "sha256")
	assert.Equal(t, nil, err, "they should be equal")

	cmd := exec.Command("/bin/sh", "-c", `echo -n huangjian|openssl sha256`)
	output, err := cmd.Output()
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, true, strings.Contains(string(output), digest), "they should be equal")
}

func TestGetDataDigestSha512(t *testing.T) {
	digest, err := GetDataDigest([]byte("huangjian"), "sha512")
	assert.Equal(t, nil, err, "they should be equal")

	cmd := exec.Command("/bin/sh", "-c", `echo -n huangjian|openssl sha512`)
	output, err := cmd.Output()
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, true, strings.Contains(string(output), digest), "they should be equal")
}

func TestGetDataDigestInvalid(t *testing.T) {
	_, err := GetDataDigest([]byte("huangjian"), "sha512aaaaaa")
	assert.NotEqual(t, nil, err, "they should be equal")
}

func TestGetDataDigestSha3_512(t *testing.T) {
	digest, err := GetDataDigest([]byte("huangjian"), "sha512")
	assert.Equal(t, nil, err, "they should be equal")

	log.Println(digest)
}
