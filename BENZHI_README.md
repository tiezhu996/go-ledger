# go-ledger

## Build, run, and test

```bash
go build ./...
go run ./cmd/ledger
go test ./...
```

This repository exposes a one-shot CLI entrypoint. A successful run prints its readiness message and exits with status 0; it is not a long-running HTTP service.

## Docker verification

```bash
docker build -f benzhi.Dockerfile -t go-ledger:local .
docker run --rm go-ledger:local
```

- Base image: `golang:1.22`
- Source directory in the image: `/app`
- Container command: `/usr/local/bin/ledger`
