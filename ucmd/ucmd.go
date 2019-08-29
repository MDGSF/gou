package ucmd

import (
	"os/exec"
)

// CmdCopy 用 cp 命令把 oldpath 拷贝到 newpath
func CmdCopy(newpath, oldpath string) error {
	cmd := exec.Command("bash", "-c", "cp -arf "+oldpath+" "+newpath)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
