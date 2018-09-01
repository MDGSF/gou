package x

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {

	key := []byte("sfe023f_9fd&fwfl")
	src := []byte("polaris@studygolang")

	result, err := AesEncrypt(src, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))

	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))

	assert.Equal(t, len(src), len(origData), "they should be the same")
	assert.Equal(t, src, origData, "they should be the same")
}
