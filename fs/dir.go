// Package fss
// File system，包括对文件、文件夹的操作。叫fss是为了防止跟自带的fs混淆
package fss

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
)

// WalkDir 遍历文件夹，包括其下的子文件夹
// @param dirPath 文件夹路径，可以是绝对路径，也可以是相对路径
// @return allSubs	字符串切片，包括子文件夹和子文件，切片里存的是相对路径还是绝对路径取决于传进来的dir是相对的还是绝对的。
//					如果需要获取到绝对路径，可以用 absPath, _ := filepath.Abs(path)
//	for example:
//		D:\MyProjects\golib\fs
//		D:\MyProjects\golib\fs\dir.go
//		D:\MyProjects\golib\fs\dir_test.go
// @return err
func WalkDir(dirPath string) (allSubs []string, err error) {
	allSubs = make([]string, 0)
	wErr := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		// 遍历某个文件/文件夹出现错误时跳过该文件/文件夹
		if err != nil {
			fmt.Printf("skipping %q: %v\n", path, errors.Unwrap(err))
			return nil
		}
		// 如果只需要返回文件，不需要文件夹，可以只加一个判断：if d.Type().IsRegular()
		_, infoErr := d.Info()
		if infoErr != nil {
			fmt.Printf("couldn't get metadata of %q: %v\n", path, infoErr)
			return nil
		}
		allSubs = append(allSubs, path)
		return nil
	})
	if wErr != nil {
		allSubs = nil
		err = fmt.Errorf("couldn't scan directory %s: %v", dirPath, wErr)
	}
	return
}
