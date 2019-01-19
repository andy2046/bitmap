#!/usr/bin/env bash

set -euo pipefail

godoc2md github.com/andy2046/bitmap \
    > $GOPATH/src/github.com/andy2046/bitmap/README.md
