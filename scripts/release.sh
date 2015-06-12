#!/bin/bash
RELEASE="$1"
echo "Creating release $1..."

go build

go get golang.org/x/tools/cmd/cover
go test -v -cover

tar -cvzf "antibody-$RELEASE-$(uname -s).tar.gz" antibody antibody.zsh
git tag "$RELEASE"
git push origin "$RELEASE"
