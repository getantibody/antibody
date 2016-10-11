package shell

import (
	"fmt"

	"github.com/kardianos/osext"
)

const template = `#!/usr/bin/env zsh
ANTIBODY_BINARY="%s"
antibody() {
	case "$1" in
	bundle)
		source <( $ANTIBODY_BINARY $@ ) 2> /dev/null || $ANTIBODY_BINARY $@
		;;
	*)
		$ANTIBODY_BINARY $@
		;;
	esac
}

_antibody() {
	IFS=' ' read -A reply <<< "$(echo "bundle update list home init help")"
}
compctl -K _antibody antibody
`

// Init returns the shell that should be loaded to antibody to work correctly.
func Init() string {
	executable, _ := osext.Executable()
	return fmt.Sprintf(template, executable)
}
