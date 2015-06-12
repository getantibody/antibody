package antibody

import (
	"github.com/caarlos0/antibody/lib/doubles"
	"testing"
)

func TestClonesValidRepo(t *testing.T) {
	home := doubles.TempHome()
	bundle := NewGitBundle("caarlos0/zsh-pg", home)
	err := bundle.Download()
	expected := home + "caarlos0-zsh-pg"
	if bundle.Folder() != expected {
		t.Error("Got", folder, "expected", expected)
	}
	if err != nil {
		t.Error("No errors expected")
	}
	assertBundledPlugins(t, 1, home)
}

func TestClonesValidRepoTwoTimes(t *testing.T) {
	home := doubles.TempHome()
	bundle := NewGitBundle("caarlos0/zsh-pg", home)
	bundle.Download()
	err := bundle.Download()
	expected := home + "caarlos0-zsh-pg"
	if bundle.Folder() != expected {
		t.Error("Got", folder, "expected", expected)
	}
	if err != nil {
		t.Error("No errors expected")
	}
	assertBundledPlugins(t, 1, home)
}

func TestClonesInvalidRepo(t *testing.T) {
	home := doubles.TempHome()
	err := NewGitBundle("this-doesnt-exist", home).Download()
	if err == nil {
		t.Error("Expected an error because this repo doesn't exist")
	}
}

func TestPullsRepo(t *testing.T) {
	home := doubles.TempHome()
	bundle := NewGitBundle("caarlos0/zsh-pg", home)
	bundle.Download()
	err := bundle.Update()
	if err != nil {
		t.Error("No errors expected")
	}
}
