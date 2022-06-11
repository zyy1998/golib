package fss

import (
	"fmt"
	"testing"
)

func TestGetUserProfile(t *testing.T) {
	userPath, _ := GetUserProfile()
	fmt.Println(userPath)
}
