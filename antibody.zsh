#!/bin/zsh
ANTIBODY_HOME="$(dirname $0)"

mkdir -p "$HOME/.antibody" || true

antibody() {
  # if [ -f "$1" ]; then
  #   cat "$1" | while read plugin; do
  #     source "$(${ANTIBODY_HOME}/antibody $plugin)"/*.plugin.zsh
  #   done
  # else
    source "$(${ANTIBODY_HOME}/antibody $*)"/*.plugin.zsh
  # fi
}
