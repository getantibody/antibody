package antibody

import (
	"os"
	"strings"
	"testing"
)

func Test_loadsDefaultHome(t *testing.T) {
	if !strings.HasSuffix(Home(), "/.antibody/") {
		t.Error()
	}
}

func Test_loadsCustomHome(t *testing.T) {
	home := "/tmp/blah"
	os.Setenv("ANTIBODY_HOME", home)
	if home != Home() {
		t.Error()
	}
}
