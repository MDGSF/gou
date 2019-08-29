package uhash

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA1(t *testing.T) {
	digest, err := SHA1([]byte("huangjian"))
	assert.Equal(t, nil, err, "they should be equal")

	cmd := exec.Command("/bin/sh", "-c", `echo -n huangjian|openssl sha1`)
	output, err := cmd.Output()
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, true, strings.Contains(string(output), digest), "they should be equal")
}
