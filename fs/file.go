package fss

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// GetFileExt gets extension of file, in lower case
// 只是获取到后缀，而不考虑该文件是否存在
func GetFileExt(path string) string {
	ext := filepath.Ext(path)
	return strings.ToLower(ext)
}

// GetBaseName 获取文件(夹)的基本名称,不考虑文件(夹)是否存在
// @param _path
// @return baseName
func GetBaseName(_path string) (baseName string) {
	return filepath.Base(_path)
}

// IsPathExist 检测文件是否存在系统中
// @param path 文件路径，可以是绝对路径，也可以是相对路径
// @return result true:存在 false:不存在
func IsPathExists(path string) (result bool) {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
