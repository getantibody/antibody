#!/bin/bash
RELEASE="v$1"
CURRENT="$(git tag | tail -n1)"
echo "Creating release $1..."

#go get golang.org/x/tools/cmd/cover
go test -v -cover ./lib

# go get github.com/mitchellh/gox
# gox -build-toolchain
gox -osarch="linux/amd64 darwin/amd64" ./...

git tag "$RELEASE"
git push origin "$RELEASE"

LOG="$(git log --pretty=oneline --abbrev-commit "$CURRENT"..HEAD)"

# go get github.com/aktau/github-release
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
    "antibody_${platform_lower}_amd64" antibody.zsh
  github-release upload \
    --user caarlos0 \
    --repo antibody \
    --tag "$RELEASE" \
    --name "$filename" \
    --file "$filename"
done
