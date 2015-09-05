#!/usr/bin/env zsh
ANTIBODY_BINARIES="$(dirname $0)"

antibody() {
  case "$1" in
  bundle|update)
    while read bundle; do
      source "$bundle" 2&> /tmp/antibody-log
    done < <( ${ANTIBODY_BINARIES}/bin/antibody $@ )
    ;;
  *)
    ${ANTIBODY_BINARIES}/bin/antibody $@
    ;;
  esac
}

_antibody() {
  IFS=' ' read -A reply <<< "$(echo "bundle update list help")"
}
compctl -K _antibody antibody
