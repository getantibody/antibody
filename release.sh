#!/bin/bash
RELEASE="$1"
echo "Creating release $1..."

go build

go get golang.org/x/tools/cmd/cover
go test -cover

tar -cvzf "antibody-$RELEASE.tar.gz" antibody antibody.zsh
git tag "$RELEASE"
git push origin "$RELEASE"
