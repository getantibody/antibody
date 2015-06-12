package antibody

import "testing"

func TestClonesValidRepo(t *testing.T) {
	home := home()
	folder, err := NewGithubBundle("caarlos0/zsh-pg", home).Download()
	expected := home + "caarlos0-zsh-pg"
	if folder != expected {
		t.Error("Got", folder, "expected", expected)
	}
	if err != nil {
		t.Error("No errors expected")
	}
	assertBundledPlugins(t, 1, home)
}

func TestClonesValidRepoTwoTimes(t *testing.T) {
	home := home()
	bundle := NewGithubBundle("caarlos0/zsh-pg", home)
	bundle.Download()
	folder, err := bundle.Download()
	expected := home + "caarlos0-zsh-pg"
	if folder != expected {
		t.Error("Got", folder, "expected", expected)
	}
	if err != nil {
		t.Error("No errors expected")
	}
	assertBundledPlugins(t, 1, home)
}

func TestClonesInvalidRepo(t *testing.T) {
	home := home()
	_, err := NewGithubBundle("this-doesnt-exist", home).Download()
	if err == nil {
		t.Error("Expected an error hence this repo doesn't exist")
	}
}

func TestPullsRepo(t *testing.T) {
	home := home()
	bundle := "caarlos0/zsh-pg"
	ghBundle := NewGithubBundle(bundle, home)
	ghBundle.Download()
	_, err := ghBundle.Update()
	if err != nil {
		t.Error("No errors expected")
	}
}
