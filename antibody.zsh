#!/usr/bin/env zsh
ANTIBODY_BINARIES="$(dirname $0)"
OPERATING_SYSTEM="$(uname -s)"
ARCHITECTURE="$(uname -m)"

antibody() {
  case "$1" in
  bundle|update)
    while read bundle; do
      source "$bundle" 2&> /tmp/antibody-log
    done < <( "${ANTIBODY_BINARIES}/bin/antibody-${OPERATING_SYSTEM}-${ARCHITECTURE}" $@ )
    ;;
  *)
    "${ANTIBODY_BINARIES}/bin/antibody-${OPERATING_SYSTEM}-${ARCHITECTURE}" $@
    ;;
  esac
}

_antibody() {
  IFS=' ' read -A reply <<< "$(echo "bundle update list help")"
}
compctl -K _antibody antibody
