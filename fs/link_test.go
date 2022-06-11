package fss

import (
	"errors"
	"os"
	"testing"
)

func TestCreateSymbolicLink(t *testing.T) {
	// os.Remove只会删除非空文件夹
	_ = os.RemoveAll("../temp")
	_ = os.Mkdir("../temp", 0750)
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name        string
		args        args
		wantSuccess bool
		wantErr     error
	}{
		{name: "创建文件符号链接", args: args{from: "D:\\MyProjects\\golib\\fs\\link_test.go", to: "../temp"}, wantSuccess: true, wantErr: nil},
		{name: "相对路径创建文件符号链接", args: args{from: "./link.go", to: "../temp"}, wantSuccess: true, wantErr: nil},
		{name: "创建文件夹符号链接", args: args{from: "D:\\MyProjects\\golib\\fs", to: "../temp"}, wantSuccess: true, wantErr: nil},
		{name: "重复创建文件符号链接", args: args{from: "D:\\MyProjects\\golib\\fs\\link_test.go", to: "../temp"}, wantSuccess: false, wantErr: errors.New("文件已存在")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuccess, err := CreateSymbolicLink(tt.args.from, tt.args.to)
			// 这里在测试error时仅简单测试是否存在
			if (tt.wantErr == nil && err != nil) || (tt.wantErr != nil && err == nil) {
				t.Errorf("CreateSymbolicLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSuccess != tt.wantSuccess {
				t.Errorf("CreateSymbolicLink() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}
