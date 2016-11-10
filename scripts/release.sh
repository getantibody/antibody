#!/bin/bash
wget -O - https://raw.githubusercontent.com/caarlos0/go-releaser/master/release |
  bash -s -- -u getantibody -r antibody -b antibody -m ./cmd/antibody/
