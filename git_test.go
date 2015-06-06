package antibody

import (
	"testing"
)

func TestClonesValidRepo(t *testing.T) {
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

func TestClonesInvalidRepo(t *testing.T) {
	home := home()
	_, err := Clone("this-doesnt-exist", home)
	if err == nil {
		t.Error("Expected an error hence this repo doesn't exist")
	}
}

func TestPullsRepo(t *testing.T) {
	home := home()
	bundle := "caarlos0/zsh-pg"
	Clone(bundle, home)
	_, err := Pull(bundle, home)
	if err != nil {
		t.Error("No errors expected")
	}
}

func TestUpdatesListOfRepos(t *testing.T) {
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
