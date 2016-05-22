#!/bin/bash

open_pdfs() {
  for i in cpu mem block; do
    go tool pprof --pdf antibody "$i.pprof" > "$1_$i.pdf"
    open "$1_$i.pdf"
  done
}

git apply "./scripts/profiling/patch.patch"
go build -o antibody ./cmd/antibody
export ANTIBODY_HOME="$(mktemp -d)"
# bundle all plugins from awesome-zsh-plugins
./antibody bundle < ./scripts/profiling/bundles.txt
open_pdfs bundle
./antibody update
open_pdfs update
./antibody list
open_pdfs list
./antibody home
open_pdfs home
./antibody init
open_pdfs init

rm -f ./antibody
git checkout ./cmd/antibody/main.go
