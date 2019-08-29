package uhash

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	digest, err := MD5([]byte("huangjian"))
	assert.Equal(t, nil, err, "they should be equal")

	cmd := exec.Command("/bin/sh", "-c", `echo -n huangjian|openssl md5`)
	output, err := cmd.Output()
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, true, strings.Contains(string(output), digest), "they should be equal")
}