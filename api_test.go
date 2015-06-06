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

func TestProcessesArgsBunde(t *testing.T) {
	home := home()
	ProcessArgs([]string{"bundle", "caarlos0/zsh-pg"}, home)
}

func TestUpdate(t *testing.T) {
	home := home()
	ProcessArgs([]string{"update"}, home)
}

func TestBundlesSinglePlugin(t *testing.T) {
	home := home()
	Bundle("caarlos0/zsh-pg", home)
}

func TestLoadsDefaultHome(t *testing.T) {
	os.Unsetenv("ANTIBODY_HOME")
	if !strings.HasSuffix(Home(), "/.antibody/") {
		t.Error("Expected default ANTIBODY_HOME")
	}
}

func TestLoadsCustomHome(t *testing.T) {
	home := home()
	if home != Home() {
		t.Error("Expected custom ANTIBODY_HOME")
	}
}

func TestFailsToBundleInvalidRepos(t *testing.T) {
	home := home()
	defer func() {
		if err := recover(); err != nil {
			t.Log("Recovered from expected error")
		} else {
			t.Error("Expected a panic hence an invalid bundle was passed")
		}
	}()
	Bundle("csadsadp", home)
}
