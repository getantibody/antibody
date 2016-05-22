#!/bin/bash

open_pdfs() {
  for i in cpu mem block; do
    go tool pprof --pdf antibody "$i.pprof" > "$1_$i.pdf" && open "$1_$i.pdf"
  done
}

git apply "./scripts/profiling/patch.patch"
go build -o antibody ./cmd/antibody
export ANTIBODY_HOME="$(mktemp -d)"
# bundle all plugins from awesome-zsh-plugins
/usr/bin/time ./antibody bundle < ./scripts/profiling/bundles.txt > /dev/null
open_pdfs bundle_download
/usr/bin/time ./antibody bundle < ./scripts/profiling/bundles.txt > /dev/null
open_pdfs bundle
/usr/bin/time ./antibody update > /dev/null
open_pdfs update
/usr/bin/time ./antibody list > /dev/null
open_pdfs list
/usr/bin/time ./antibody home > /dev/null
open_pdfs home
/usr/bin/time ./antibody init > /dev/null
open_pdfs init

rm -f ./antibody
git checkout ./cmd/antibody/main.go
