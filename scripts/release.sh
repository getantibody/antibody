#!/bin/bash
CURRENT="$(git describe --tags --abbrev=0)"
PREVIOUS=$(git describe --tags --abbrev=0 ${CURRENT}^)

echo "Installing needed tools..."
go get github.com/mitchellh/gox
gox -build-toolchain
go get github.com/aktau/github-release
go get golang.org/x/tools/cmd/cover

echo "Creating release $CURRENT..."
gox \
  -output="./bin/{{.Dir}}_{{.OS}}_{{.Arch}}" \
  -os="linux darwin freebsd openbsd netbsd" \
  -ldflags="-X main.version $CURRENT" \
  ./cmd/antibody/
LOG="$(git log --pretty=oneline --abbrev-commit "$PREVIOUS".."$CURRENT")"
github-release release \
  --user caarlos0 \
  --repo antibody \
  --tag "$CURRENT" \
  --name "$2" \
  --description "$LOG" \
  --pre-release
if [ "$(uname -s)" = "Darwin" ]; then
  PATH="/usr/local/opt/gnu-tar/libexec/gnubin:$PATH"
fi
# shellcheck disable=SC2012
ls ./bin | while read file; do
  filename="$file.tar.gz"
  tar \
    --transform="s/${file}/antibody/" \
    -cvzf "$filename" \
    "bin/${file}" antibody.zsh README.md LICENSE
  github-release upload \
    --user caarlos0 \
    --repo antibody \
    --tag "$CURRENT" \
    --name "$filename" \
    --file "$filename"
done
