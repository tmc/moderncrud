# syntax=docker/dockerfile:1.3

FROM golang:1.17

WORKDIR /go/src/github.com/tmc/moderncrud
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg go mod download -x
RUN --mount=type=cache,target=/go/pkg --mount=type=cache,target=/root/.cache/go-build go get github.com/mattn/go-sqlite3
RUN --mount=type=cache,target=/go/pkg --mount=type=cache,target=/root/.cache/go-build go install github.com/benbjohnson/litestream/cmd/litestream

COPY . .
RUN --mount=type=cache,target=/go/pkg --mount=type=cache,target=/root/.cache/go-build go install -v ./cmd/...

CMD ["./scripts/entrypoint.sh"]

