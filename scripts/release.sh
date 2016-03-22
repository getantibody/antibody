#!/bin/bash
export GO15VENDOREXPERIMENT=1

CURRENT="$(git describe --tags --abbrev=0)"
PREVIOUS=$(git describe --tags --abbrev=0 ${CURRENT}^)

echo "Installing needed tools..."
go get github.com/mitchellh/gox
go get github.com/aktau/github-release

echo "Creating release $CURRENT..."
gox \
  -output="./bin/{{.Dir}}_{{.OS}}_{{.Arch}}" \
  -os="linux darwin freebsd openbsd netbsd" \
  -ldflags="-X main.version $CURRENT" \
  ./cmd/antibody/
LOG="$(git log --pretty=oneline --abbrev-commit "$PREVIOUS".."$CURRENT")"
DESCRIPTION="$LOG\n\nCompiled with: $(go version)"
github-release release \
  --user getantibody \
  --repo antibody \
  --tag "$CURRENT" \
  --description "$DESCRIPTION" \
  --pre-release

for file in ./bin/*; do
  filename="$(grep "$file" ./scripts/version_map | cut -f2 -d'=').tar.gz"
  tar \
    --transform="s/${file}/antibody/" \
    -cvzf "$filename" \
    "bin/${file}" antibody.zsh README.md LICENSE
  github-release upload \
    --user getantibody \
    --repo antibody \
    --tag "$CURRENT" \
    --name "$filename" \
    --file "$filename"
done
