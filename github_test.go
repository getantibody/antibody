package antibody

import (
	"os"
	"testing"
)

func home() string {
	return os.TempDir()
}

func Test_CloneValidRepo(t *testing.T) {
	home := home()
	folder, err := Clone("caarlos0/antibody", home)
	if folder != home+"caarlos0-antibody" || err != nil {
		t.Error()
	}
}

func Test_CloneInvalidRepo(t *testing.T) {
	home := home()
	_, err := Clone("this-doesnt-exist", home)
	if err == nil {
		t.Error()
	}
}

func Test_Pull(t *testing.T) {
	home := home()
	bundle := "caarlos0/antibody"
	Clone(bundle, home)
	_, err := Pull(bundle, home)
	if err != nil {
		t.Error()
	}
}
