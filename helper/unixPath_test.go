package helper

import (
	"runtime"
	"testing"
)

func TestConvertToUnixPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		out := ConvertToUnixPath("C:/Windows/")
		want := "/c/Windows/"
		if out != want {
			t.Fatalf("test convert to unix path : incorrect result: want: %q ,got: %q", want, out)
		}
	}
}
