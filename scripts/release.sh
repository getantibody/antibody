#!/bin/sh
set -e

# enables go 1.5 vendor expirement
export GO15VENDOREXPERIMENT=1

cleanup() {
  rm -rf dist
}

# normalize Golang's OS and Arch to uname compatibles
normalize() {
  echo "$1" | sed \
    -e 's/darwin/Darwin/' \
    -e 's/linux/Linux/' \
    -e 's/freebsd/FreeBSD/' \
    -e 's/openbsd/OpenBSD/' \
    -e 's/netbsd/NetBSD/' \
    -e 's/386/i386/' \
    -e 's/amd64/x86_64/'
}

# builds the binaries with gox
build() {
  echo "Building $CURRENT..."
  go get github.com/mitchellh/gox
  gox -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}/antibody" \
    -os="linux darwin freebsd openbsd netbsd" \
    -ldflags="-X main.version $CURRENT" \
    ./cmd/antibody/
}

# package the binaries in .tar.gz files
package() {
  echo "Packaging $CURRENT..."
  for folder in ./dist/*; do
    local filename="$(normalize "$folder").tar.gz"
    tar -cvzf "$filename" "${folder}/antibody" antibody.zsh README.md LICENSE
  done
}

# release it to github
release() {
  echo "Releasing $CURRENT..."
  local -r log="$(git log --pretty=oneline --abbrev-commit "$PREVIOUS".."$CURRENT")"
  local -r description="${log}\n\nBuilt with: $(go version)"
  go get github.com/aktau/github-release
  echo "Creating release $CURRENT..."
  github-release release \
    --user getantibody \
    --repo antibody \
    --tag "$CURRENT" \
    --description "$description" \
    --pre-release ||
    github-release edit \
      --user getantibody \
      --repo antibody \
      --tag "$CURRENT" \
      --description "$description" \
      --pre-release
  for file in ./dist/*.tar.gz; do
    echo "--> Uploading $file..."
    github-release upload \
      --user getantibody \
      --repo antibody \
      --tag "$CURRENT" \
      --name "$(echo $file | sed 's/\.\/dist\///')" \
      --file "$file"
  done
}

cleanup
CURRENT="$(git describe --tags --abbrev=0)"
PREVIOUS=$(git describe --tags --abbrev=0 ${CURRENT}^)
build
package
release
