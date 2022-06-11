package array

import (
	"fmt"
	"testing"
)

func Test2DimensionalArray(t *testing.T) {
	arr := [][3]string{
		{"I", "love", "you"},
		{"You", "love", "me"},
	}
	fmt.Printf("arr value:%#v type:%T\n", arr, arr)
	for _, ele := range arr {
		for _, item := range ele {
			fmt.Println(item)
		}
	}
}
