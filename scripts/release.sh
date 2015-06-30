#!/bin/bash
RELEASE="v$1"
CURRENT="$(git tag | tail -n1)"

# echo "Installing needed tools..."
# go get github.com/mitchellh/gox
# gox -build-toolchain
# go get github.com/aktau/github-release
# go get golang.org/x/tools/cmd/cover

echo "Creating release $1..."
go test -v -cover
rm -rf ./bin/
rm -rf ./*.tar.gz
gox \
  -output="./bin/{{.Dir}}_{{.OS}}_{{.Arch}}" \
  -os="linux darwin freebsd openbsd netbsd" \
  ./...
git tag "$RELEASE"
git push origin "$RELEASE"
sed -i'' "s/HEAD/$RELEASE/g" antibody.zsh
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
# shellcheck disable=SC2012
ls ./bin | while read file; do
  filename="$file.tar.gz"
  tar \
    --transform="s/${file}/antibody/" \
    -cvzf "$filename" \
    "bin/${file}" antibody.zsh
  github-release upload \
    --user caarlos0 \
    --repo antibody \
    --tag "$RELEASE" \
    --name "$filename" \
    --file "$filename"
done
rm -rf ./*.tar.gz
