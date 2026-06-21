# Mehrfunds Node

**mehrd** is the Mehrfunds blockchain node, built with Cosmos SDK and CometBFT.

## Get started

```
ignite chain serve
```

`serve` installs dependencies, builds, initializes, and starts the chain in development mode.

### Configure

The chain can be configured with `config.yml`. See the [Ignite CLI docs](https://docs.ignite.com) for details.

## Release

Tag a release with a `v` prefix to trigger a build:

```
git tag v0.1
git push origin v0.1
```

### Install

```
curl https://get.ignite.com/mehrfunds/node@latest! | sudo bash
```

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [CometBFT](https://docs.cometbft.com)
