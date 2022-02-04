package ghelp

import (
	"io"
	"os"
)

// IsFile 是否文件
func IsFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	if f.IsDir() {
		return false
	}
	return true
}

// IsDir 是否文件夹
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	if f.IsDir() {
		return true
	}
	return false
}

// CopyFile 拷贝文件
func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	defer src.Close()
	if err != nil {
		return
	}
	srcF, err := src.Stat()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, srcF.Mode())
	defer dst.Close()
	if err != nil {
		return
	}
	return io.Copy(dst, src)
}
