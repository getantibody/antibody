#!/bin/bash
rm -rf *.pdf
rm -rf *.pprof
current="$(dirname $0)"
git apply "$current/patch.patch"
go build -o antibody ./cmd/antibody
export ANTIBODY_HOME="$(mktemp -d)"
./antibody bundle < ./scripts/profiling/bundles.txt
for i in cpu mem block; do
  go tool pprof --pdf antibody "$i.pprof" > "$i.pdf"
  open "$i.pdf"
done
rm -rf ./antibody
git checkout ./cmd/antibody/main.go
