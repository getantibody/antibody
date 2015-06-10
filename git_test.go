package main

import "testing"

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
	assertBundledPlugins(t, 1, home)
}

func TestClonesValidRepoTwoTimes(t *testing.T) {
	home := home()
	Clone("caarlos0/zsh-pg", home)
	folder, err := Clone("caarlos0/zsh-pg", home)
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
