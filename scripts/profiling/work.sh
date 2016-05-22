#!/bin/bash
git apply "./scripts/profiling/patch.patch"
go build -o antibody ./cmd/antibody
export ANTIBODY_HOME="$(mktemp -d)"
# bundle all plugins from awesome-zsh-plugins
./antibody bundle < ./scripts/profiling/bundles.txt
for i in cpu mem block; do
  go tool pprof --pdf antibody "$i.pprof" > "$i.pdf"
  open "$i.pdf"
done
rm -f ./antibody
git checkout ./cmd/antibody/main.go
