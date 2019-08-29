package ucmd

import (
	"os"
	"testing"

	"github.com/MDGSF/utils"
)

func TestExecCopy1(t *testing.T) {
	os.MkdirAll("/tmp/anno_test_copy1", 0777)
	CmdCopy("/tmp/anno_test_copy2", "/tmp/anno_test_copy1")

	ret := utils.IsDir("/tmp/anno_test_copy1")
	if !ret {
		t.Fatal(ret)
	}

	ret = utils.IsDir("/tmp/anno_test_copy2")
	if !ret {
		t.Fatal(ret)
	}

	os.RemoveAll("/tmp/anno_test_copy1")
	os.RemoveAll("/tmp/anno_test_copy2")
}

func TestExecCopy2(t *testing.T) {
	err := CmdCopy("", "/tmp/anno_test_copy1")
	if err == nil {
		t.Fatal(err)
	}
}
