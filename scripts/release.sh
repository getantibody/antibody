#!/bin/bash
export GO15VENDOREXPERIMENT=1

CURRENT="$(git describe --tags --abbrev=0)"
PREVIOUS=$(git describe --tags --abbrev=0 ${CURRENT}^)

echo "Installing needed tools..."
go get github.com/mitchellh/gox
gox -build-toolchain
go get github.com/aktau/github-release
go get golang.org/x/tools/cmd/cover

declare -A gox_to_uname
gox_to_uname=(
	[antibody_darwin_386]='antibody-Darwin-i386'
	[antibody_darwin_amd64]='antibody-Darwin-x86_64'
	[antibody_darwin_arm]='antibody-Darwin-arm'

	[antibody_linux_386]='antibody-Linux-i386'
	[antibody_linux_amd64]='antibody-Linux-x86_64'
	[antibody_linux_arm]='antibody-Linux-arm'

	[antibody_freebsd_386]='antibody-FreeBSD-i386'
	[antibody_freebsd_amd64]='antibody-FreeBSD-x86_64'
	[antibody_freebsd_arm]='antibody-FreeBSD-arm'

	[antibody_openbsd_386]='antibody-OpenBSD-i386'
	[antibody_openbsd_amd64]='antibody-OpenBSD-x86_64'
	[antibody_openbsd_arm]='antibody-OpenBSD-arm'

	[antibody_netbsd_386]='antibody-NetBSD-i386'
	[antibody_netbsd_amd64]='antibody-NetBSD-x86_64'
	[antibody_netbsd_arm]='antibody-NetBSD-arm'
)


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
# shellcheck disable=SC2012
ls ./bin | while read file; do
  filename="${gox_to_uname[$file]}.tar.gz"
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
