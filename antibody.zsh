#!/bin/zsh
ANTIBODY_BINARIES="$(dirname $0)"
mkdir -p "$HOME/.antibody" || true

antibody() {
  while read bundle; do
    source "$bundle"/*.plugin.zsh &
  done < <( ${ANTIBODY_BINARIES}/antibody $@ )
}
