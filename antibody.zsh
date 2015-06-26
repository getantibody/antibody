#!/usr/bin/env zsh
ANTIBODY_BINARIES="$(dirname $0)"
[[ -n "${ANTIBODY_HOME}" ]] && mkdir -p "${ANTIBODY_HOME}" || mkdir -p "${HOME}/.antibody"

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
