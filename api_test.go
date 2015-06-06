package main

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

func expectError(t *testing.T) {
	if err := recover(); err != nil {
		t.Log("Recovered from expected error")
	} else {
		t.Error("Expected an error here!")
	}
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
	defer expectError(t)
	Bundle("csadsadp", home)
}

func TestFailsToProcessInvalidArgs(t *testing.T) {
	home := home()
	defer expectError(t)
	ProcessArgs([]string{"nope", "caarlos0/zsh-pg"}, home)
}

func TestReadsStdinIsFalse(t *testing.T) {
	if ReadStdin() {
		t.Error("Not reading STDIN")
	}
}

func TestReadsStdinIsTrue(t *testing.T) {
	os.Stdin.Write([]byte("Some STDIN"))
	if ReadStdin() {
		t.Error("Not reading STDIN")
	}
}
