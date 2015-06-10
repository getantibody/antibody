package main

import (
	"os"
	"os/exec"
	"strings"
)

const GH = "https://github.com/"

func folder(bundle string, home string) string {
	return home + strings.Replace(bundle, "/", "-", -1)
}

func Clone(bundle string, home string) (string, error) {
	folder := folder(bundle, home)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		clone := exec.Command("git", "clone", "--depth", "1", GH+bundle, folder)
		return folder, clone.Run()
	}
	return folder, nil
}

func Pull(bundle string, home string) (string, error) {
	folder := folder(bundle, home)
	pull := exec.Command("git", "-C", folder, "pull", "origin", "master")
	return folder, pull.Run()
}
