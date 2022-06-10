package fss

import "testing"

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
