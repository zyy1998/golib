package fss

import (
	"testing"
)

func TestGetFileExt(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "t1", args: args{"file_test.go"}, want: ".go"},
		{name: "t2", args: args{"../../READE.md"}, want: ".md"},
		{name: "t3", args: args{"xxx"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileExt(tt.args.path); got != tt.want {
				t.Errorf("GetFileExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPathExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{name: "t1", args: args{path: "./filex.go"}, wantResult: false},
		{name: "t2", args: args{path: "./file.go"}, wantResult: true},
		{name: "t3", args: args{path: "../fs"}, wantResult: true},
		//下面测试仅能在windows进行
		{name: "t10", args: args{path: "c:/users"}, wantResult: true},
		{name: "t11", args: args{path: "c:/usersx"}, wantResult: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := IsPathExists(tt.args.path); gotResult != tt.wantResult {
				t.Errorf("IsFileExists() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetBaseName(t *testing.T) {
	type args struct {
		_path string
	}
	tests := []struct {
		name         string
		args         args
		wantBaseName string
	}{
		{name: "t1", args: args{_path: "../sync/wait_group_test.go"}, wantBaseName: "wait_group_test.go"},
		{name: "t2", args: args{_path: "D:\\MyProjects\\golib\\fs\\link_test.go"}, wantBaseName: "link_test.go"},
		{name: "t3", args: args{_path: "../fs"}, wantBaseName: "fs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBaseName := GetBaseName(tt.args._path); gotBaseName != tt.wantBaseName {
				t.Errorf("GetBaseName() = %v, want %v", gotBaseName, tt.wantBaseName)
			}
		})
	}
}
