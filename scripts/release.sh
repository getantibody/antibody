#!/bin/bash
RELEASE="v$1"
CURRENT="$(git tag | tail -n1)"

# echo "Installing needed tools..."
# go get github.com/mitchellh/gox
# gox -build-toolchain
# go get github.com/aktau/github-release
# go get golang.org/x/tools/cmd/cover

echo "Creating release $1..."
go test -v -cover ./lib
gox \
  -output="./bin/{{.Dir}}_{{.OS}}_{{.Arch}}" \
  -osarch="linux/amd64 darwin/amd64" \
  ./...
git tag "$RELEASE"
git push origin "$RELEASE"
LOG="$(git log --pretty=oneline --abbrev-commit "$CURRENT"..HEAD)"
github-release release \
  --user caarlos0 \
  --repo antibody \
  --tag "$RELEASE" \
  --name "$2" \
  --description "$LOG" \
  --pre-release
if [ "$(uname -s)" = "Darwin" ]; then
  PATH="/usr/local/opt/gnu-tar/libexec/gnubin:$PATH"
fi
for platform in Darwin Linux; do
  filename="antibody-$RELEASE-$platform.tar.gz"
  platform_lower="$(echo $platform | tr '[:upper:]' '[:lower:]')"
  tar \
    --transform="s/_${platform_lower}_amd64//" \
    -cvzf "$filename" \
    "bin/antibody_${platform_lower}_amd64" antibody.zsh
  github-release upload \
    --user caarlos0 \
    --repo antibody \
    --tag "$RELEASE" \
    --name "$filename" \
    --file "$filename"
done
