package shell

import (
	"fmt"

	"github.com/kardianos/osext"
)

const template = `#!/usr/bin/env zsh
ANTIBODY_BINARY="%s"
antibody() {
	case "$1" in
	bundle|update)
		while read bundle; do
			source "$bundle" 2&> /tmp/antibody-log
		done < <( $ANTIBODY_BINARY $@ )
		;;
	*)
		$ANTIBODY_BINARY $@
		;;
	esac
}

_antibody() {
	IFS=' ' read -A reply <<< "$(echo "bundle update list help")"
}
compctl -K _antibody antibody
`

// Init returns the shell that should be loaded to antibody to work correctly.
func Init() (string, error) {
	executable, err := osext.Executable()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(template, executable), nil
}
