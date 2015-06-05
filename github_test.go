package antibody

import (
	"io/ioutil"
	"os"
	"testing"
)

func home() string {
	file, err := ioutil.TempDir(os.TempDir(), "antibody")
	if err != nil {
		panic(err.Error())
	}
	defer os.RemoveAll(file)
	return file + "/"
}

func Test_ClonesValidRepo(t *testing.T) {
	home := home()
	folder, err := Clone("caarlos0/zsh-pg", home)
	expected := home + "caarlos0-zsh-pg"
	if folder != expected {
		t.Error("Got", folder, "expected", expected)
	}
	if err != nil {
		t.Error("No errors expected")
	}
}

func Test_ClonesInvalidRepo(t *testing.T) {
	home := home()
	_, err := Clone("this-doesnt-exist", home)
	if err == nil {
		t.Error("Expected an error hence this repo doesn't exist")
	}
}

func Test_PullsRepo(t *testing.T) {
	home := home()
	bundle := "caarlos0/zsh-pg"
	Clone(bundle, home)
	_, err := Pull(bundle, home)
	if err != nil {
		t.Error("No errors expected")
	}
}

func Test_UpdatesListOfRepos(t *testing.T) {
	home := home()
	bundle1 := "caarlos0/zsh-pg"
	bundle2 := "caarlos0/zsh-add-upstream"
	Clone(bundle1, home)
	Clone(bundle2, home)
	bundles, err := Update(home)
	if err != nil {
		t.Error("No errors expected")
	}
	if len(bundles) != 2 {
		t.Error(len(bundles), "updated bundles, expected 2")
	}
}
