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
	folder := Clone("caarlos0/antibody", home)
	if folder != home+"caarlos0-antibody" {
		t.Error()
	}
}

func Test_CloneInvalidRepo(t *testing.T) {
	home := home()
	folder, err := Clone("caarlos0/asdasdasd-asdasdas-as222", home)
	if err == nil {
		t.Error()
	}
}
