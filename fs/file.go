package fss

import (
	"path/filepath"
	"strings"
)

// GetFileExt gets extension of file, in lower case
// 只是获取到后缀，而不考虑该文件是否存在
func GetFileExt(path string) string {
	ext := filepath.Ext(path)
	return strings.ToLower(ext)
}
