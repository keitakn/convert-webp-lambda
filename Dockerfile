FROM golang:1.16-alpine3.12

LABEL maintainer="https://github.com/keitakn"

WORKDIR /go/app

COPY . .

ARG GOLANGCI_LINT_VERSION=v1.38.0

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl make && \
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VERSION} && \
  go install golang.org/x/tools/cmd/goimports@latest

ENV CGO_ENABLED 0
