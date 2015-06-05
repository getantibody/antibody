package antibody

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func home() string {
	file, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	defer os.RemoveAll(file)
	os.Setenv("ANTIBODY_HOME", file+"/")
	return file + "/"
}

func Test_bundlesSinglePlugin(t *testing.T) {
	home := home()
	Bundle("caarlos0/zsh-pg", home)
}

func Test_loadsDefaultHome(t *testing.T) {
	os.Unsetenv("ANTIBODY_HOME")
	if !strings.HasSuffix(Home(), "/.antibody/") {
		t.Error("Expected default ANTIBODY_HOME")
	}
}

func Test_loadsCustomHome(t *testing.T) {
	home := home()
	if home != Home() {
		t.Error("Expected custom ANTIBODY_HOME")
	}
}
