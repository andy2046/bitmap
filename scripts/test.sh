#!/usr/bin/env bash

set -euo pipefail

go fmt .
go vet .
golint -set_exit_status .

GOCACHE=off go test -v .
