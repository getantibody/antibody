#!/bin/zsh
ANTIBODY_BINARIES="$(dirname $0)"
mkdir -p "$HOME/.antibody" || true

antibody() {
  local bundles="$(${ANTIBODY_BINARIES}/antibody $@)"
  echo $bundles | while read bundle; do
    source "$bundle"/*.plugin.zsh
  done
}
