package utils

import "os"

// IsFile check path is file or not.
func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.Mode().IsRegular()
}

// IsDir check path is directory or not.
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.Mode().IsDir()
}
