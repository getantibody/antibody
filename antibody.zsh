#!/bin/zsh
ANTIBODY_BINARIES="$(dirname $0)"
mkdir -p "$HOME/.antibody" || true

antibody() {
  while read bundle; do
    source "$bundle"/*.plugin.zsh 2&> /tmp/antibody-log
  done < <( ${ANTIBODY_BINARIES}/bin/antibody $@ )
}

_antibody() {
  IFS=' ' read -A reply <<< "$(echo "bundle update")"
}
compctl -K _antibody antibody
