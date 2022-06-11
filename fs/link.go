package fss

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// CreateSymbolicLink 创建符号链接
// windows中需要管理员权限 如果在goland中运行，可以在run/debug configuration中勾选`run with elevated privileges`
// @param from 	要给谁创建链接，可以是文件也可以是文件夹
// @param to	存储位置，是个文件夹，默认以源文件(夹)的basename创建链接
// @return success
// @return err
func CreateSymbolicLink(from string, to string) (success bool, err error) {
	if !IsPathExists(from) {
		success = false
		err = fmt.Errorf("%v dose not exist", from)
		return
	}
	if !IsPathExists(to) {
		success = false
		err = fmt.Errorf("%v dose not exist", to)
		return
	}
	// 经测试，给文件夹创建符号链接时必须使用绝对路径，所以干脆在这里都使用绝对路径
	from, _ = filepath.Abs(from)
	to, _ = filepath.Abs(to)
	baseName := filepath.Base(from)
	to = path.Join(to, baseName)
	err = os.Symlink(from, to)
	if err != nil {
		fmt.Println(err.Error())
		success = false
		return
	}
	success = true
	return
}
