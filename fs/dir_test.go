package fss

import (
	"fmt"
	"testing"
)

func TestWalkDir(t *testing.T) {
	dirPath := "../fs"
	//dirPath := "D:\\MyProjects\\golib"
	allFiles, _ := WalkDir(dirPath)
	for _, file := range allFiles {
		fmt.Printf("%v\n", file)
	}
}
