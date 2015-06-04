#!/bin/zsh
ANTIBODY_HOME="$(dirname $0)"

mkdir -p "$HOME/.antibody" || true

antibody() {
  source "$(${ANTIBODY_HOME}/antibody $@)"/*.plugin.zsh
}
