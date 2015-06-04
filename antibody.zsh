#!/bin/zsh
ANTIBODY_HOME="$(dirname $0)"

mkdir -p "$HOME/.antibody" || true

antibody() {
  local bundles="$(${ANTIBODY_HOME}/antibody $@)"
  echo $bundles | while read bundle; do
    source "$bundle"/*.plugin.zsh
  done
}
