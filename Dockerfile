FROM golang:1.17

WORKDIR /go/src/github.com/tmc/moderncrud
COPY go.mod go.sum ./
RUN go mod download -x
RUN go get github.com/mattn/go-sqlite3
RUN go install github.com/benbjohnson/litestream/cmd/litestream

COPY . .
RUN go install -v ./cmd/...

CMD ["./scripts/entrypoint.sh"]

