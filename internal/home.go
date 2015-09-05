package internal

import (
	"io/ioutil"
	"os"
)

// TempHome creates a new folder in TempDir, sets it as ANTIBODY_HOME and
// returns its fullpath.
func TempHome() string {
	file, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	os.Setenv("ANTIBODY_HOME", file+"/")
	return file + "/"
}
