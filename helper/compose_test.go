package helper

import "testing"

func TestComposeSource(t *testing.T) {
	out := ComposeSource("/path/to/script.zsh")
	want := "source /path/to/script.zsh"
	if out != want {
		t.Fatalf("test compose source: want: %q,got: %q", want, out)
	}
}
func TestComposeFPath(t *testing.T) {
	out := ComposeFPath("/path/to/fpath")
	want := "fpath+=( /path/to/fpath )"
	if out != want {
		t.Fatalf("test compose fpath: want: %q,got: %q", want, out)
	}
}

func TestComposeEnvPath(t *testing.T) {
	out := ComposeEnvPath("/path/to/bin")
	want := "export PATH=\"/path/to/bin:$PATH\""
	if out != want {
		t.Fatalf("test compose envpath: want:%q,got: %q", want, out)
	}
}
