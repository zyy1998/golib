package fss

import "os"

// GetUserProfile 获取用户home路径
// @return string
// @return error
func GetUserProfile() (string, error) {
	return os.UserHomeDir()
}
