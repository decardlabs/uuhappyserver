# uuhappyserver

uuhappyserver is the server-side core project used with uuhappy client.

It is a Go-based proxy core forked from Xray-core and rebranded for the uuhappy ecosystem.

## Repository Layout

- `main/`: CLI entrypoints and commands
- `core/`: core runtime and version/build info
- `app/`: application features
- `proxy/`: protocol implementations
- `transport/`: transport layer
- `infra/`: config and infrastructure helpers
- `deploy_uuhappyserver.sh`: example Linux deployment script

## Prerequisites

- Go 1.26+ (follow `go.mod`)
- Git
- Linux server with systemd (for script deployment)

## Build

From repository root:

```bash
CGO_ENABLED=0 go build -o uuhappyserver -trimpath -buildvcs=false -ldflags="-s -w -buildid=" -v ./main
```

Windows:

```powershell
$env:CGO_ENABLED=0
go build -o uuhappyserver.exe -trimpath -buildvcs=false -ldflags="-s -w -buildid=" -v ./main
```

## Run

```bash
./uuhappyserver run -c /path/to/config.json
```

## Quick Deploy (Linux)

A deployment helper script is provided:

```bash
chmod +x deploy_uuhappyserver.sh
sudo ./deploy_uuhappyserver.sh
```

The script will:

- create default config under `/usr/local/etc/uuhappyserver/`
- register systemd service `uuhappyserver.service`
- start and enable the service

## Test

```bash
go test ./...
```

## Related

- Client repo: `https://github.com/decardlabs/uuhappy`
- Server repo: `https://github.com/decardlabs/uuhappyserver`

## License

See `LICENSE`.
