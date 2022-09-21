#!/bin/sh

pathReturn="$(pwd)"

# Install stringer command
cd "/tmp" || exit 1
go install "golang.org/x/tools/cmd/stringer@latest"

which stringer || {
    echo "stringer command not found"
    exit 1
}

# Run stringer command
cd "${pathReturn}" || exit 1
go generate ./...
