package antibody

import (
	"os"
	"strings"
	"testing"
)

func mockHome(t *testing.T) string {
	home := os.TempDir()
	t.Log("Using home ", home)
	os.Setenv("ANTIBODY_HOME", home)
	return home
}

func Test_loadsDefaultHome(t *testing.T) {
	if !strings.HasSuffix(Home(), "/.antibody/") {
		t.Error()
	}
}

func Test_loadsCustomHome(t *testing.T) {
	home := mockHome(t)
	if home != Home() {
		t.Error()
	}
}
