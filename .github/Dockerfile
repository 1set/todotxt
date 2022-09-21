# =============================================================================
#  Test Container for Vaious Go Versions (Last Updated: 2022-09-21)
# =============================================================================
# Default version: golang:alpine
ARG VARIANT="alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

RUN apk add --no-cache \
    git \
    alpine-sdk \
    build-base

WORKDIR /workspaces

ENTRYPOINT go mod download && go test -race ./...
