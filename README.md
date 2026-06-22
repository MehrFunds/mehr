# mehr

**mehrd** is the Mehr Funds blockchain node — a sovereign Cosmos SDK chain with CometBFT consensus. It powers the on-chain state for the Mehr Funds platform: watches, webhooks, and API key registrations.

- Chain ID: `mehr-1`
- Bond denom: `umehr` (1 MEHR = 1,000,000 umehr)
- Bech32 prefix: `mehr`

## Install

```sh
curl -fsSL https://raw.githubusercontent.com/MehrFunds/mehr/refs/heads/main/install.sh | sh
```

The script detects your OS and architecture, downloads the right binary, optionally initializes a node, and drops into an interactive admin panel.

### Manual download

Pre-built binaries are attached to every [release](https://github.com/MehrFunds/mehr/releases):

| Binary | Platform |
|--------|----------|
| `mehrd-linux-amd64` | Linux x86-64 |
| `mehrd-linux-arm64` | Linux ARM64 (Graviton, RPi 4+) |
| `mehrd-linux-arm` | Linux ARMv7 (RPi 3) |
| `mehrd-darwin-amd64` | macOS Intel |
| `mehrd-darwin-arm64` | macOS Apple Silicon |
| `mehrd-windows-amd64.exe` | Windows x86-64 |
| `mehrd-windows-arm64.exe` | Windows ARM64 |

## Run a node

```sh
# Initialize (once)
mehrd init <moniker> --chain-id mehr-1

# Start
mehrd start
```

Configuration lives in `~/.mehr/config/`. The two files to know:

- `config.toml` — P2P, RPC, consensus settings
- `app.toml` — REST API, gRPC, pruning, mempool

## Release

Tag a commit to trigger the release workflow:

```sh
git tag v0.2
git push origin v0.2
```

GitHub Actions cross-compiles all 7 platform binaries and publishes them to the release page automatically.

## Development

Requires Go 1.25+.

```sh
# Build locally
go build ./cmd/mehrd

# Run tests
go test ./...
```

## Links

- [Mehr Funds](https://mehrfunds.com)
- [Dashboard](https://dash.mehrfunds.com)
- [Cosmos SDK](https://docs.cosmos.network)
- [CometBFT](https://docs.cometbft.com)
