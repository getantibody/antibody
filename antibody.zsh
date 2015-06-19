#!/bin/zsh
ANTIBODY_BINARIES="$(dirname $0)"
mkdir -p "$HOME/.antibody" || true

antibody() {
  case "$1" in
  version)
    echo "HEAD"
    ;;
  *)
    while read bundle; do
      source "$bundle"/*.plugin.zsh 2&> /tmp/antibody-log
    done < <( ${ANTIBODY_BINARIES}/bin/antibody $@ )
    ;;
  esac
}

_antibody() {
  IFS=' ' read -A reply <<< "$(echo "bundle update version")"
}
compctl -K _antibody antibody
