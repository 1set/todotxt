#!/bin/sh
# =============================================================================
#  This script updates Go modules to the latest version.
# =============================================================================
#  It will remove the go.mod file and run `go mod tidy` to get the latest moule
#  versions.
#  Then it will run the tests to make sure the code is still working, and fails
#  if any errors are found during the process.
#
#  NOTE: This script is aimed to run in the container via docker-compose.
#    See "tidy" service: ./docker-compose.yml
# =============================================================================

set -eu

echo '* Backup modules ...'
mv go.mod go.mod.bak
mv go.sum go.sum.bak

echo '* Create new blank go.mod ...'
# Copy the first 4 lines of the go.mod.bak file to the new go.mod file.
< go.mod.bak head -n 4 > go.mod

echo '* Get latest modules ...'
go get "github.com/pkg/errors"
go get "github.com/stretchr/testify"
go get "github.com/1set/gut"

echo '* Run go tidy ...'
go mod tidy

echo '* Run tests ...'
go test ./... && {
    echo '* Testing passed. Removing old go.mod file ...'
    rm -f go.mod.bak
    rm -f go.sum.bak
    echo 'Successfully updated modules!'
}