package helper

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

//ConvertToUnixPath on Windows convert path to Unix style, use cygpath.exe provided by msys2/cygwin
//On other OS do nothing
//Windows native program unable access files through converted path.
//Call this function only when intend to compose an output for other program(such as zsh) use.
func ConvertToUnixPath(path string) string {
	if runtime.GOOS == "windows" {
		out, err := exec.Command("cygpath.exe", path).Output()
		if err != nil {
			log.Fatalf("Error: Unable convert Windows path to unix: %v", err)
		}
		return strings.TrimSpace(string(out))
	}
	return path
}
