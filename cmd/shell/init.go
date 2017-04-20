package shell

import (
	"bytes"
	"os"
	"text/template"
)

const tmpl = `#!/usr/bin/env zsh
antibody() {
	case "$1" in
	bundle)
		source <( {{ .Executable }} $@ ) 2> /dev/null || {{ .Executable }} $@
		;;
	*)
		{{ .Executable }} $@
		;;
	esac
}

_antibody() {
	IFS=' ' read -A reply <<< "$(echo "bundle update list home init help")"
}
compctl -K _antibody antibody
`

// Init returns the shell that should be loaded to antibody to work correctly.
func Init() (string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", err
	}
	var template = template.Must(template.New("init").Parse(tmpl))
	var out bytes.Buffer
	err = template.Execute(&out, struct {
		Executable string
	}{
		Executable: executable,
	})
	return out.String(), err
}
