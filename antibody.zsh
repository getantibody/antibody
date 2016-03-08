#!/usr/bin/env zsh
ANTIBODY_DIRECTORY="$(dirname $0)"
OPERATING_SYSTEM="$(uname -s)"
ARCHITECTURE="$(uname -m)"

antibody() {
  antibody_os_arch=${ANTIBODY_DIRECTORY}/bin/antibody-${OPERATING_SYSTEM}-${ARCHITECTURE}

  if [[ -x "$antibody_os_arch" ]]; then
    antibody="$antibody_os_arch"
  else
    antibody="${ANTIBODY_DIRECTORY}/bin/antibody"
  fi

  case "$1" in
  bundle|update)
    while read bundle; do
      source "$bundle" 2&> /tmp/antibody-log
    done < <( "$antibody" $@ )
    ;;
  *)
    "$antibody" $@
    ;;
  esac
}

_antibody() {
  IFS=' ' read -A reply <<< "$(echo "bundle update list help")"
}
compctl -K _antibody antibody
