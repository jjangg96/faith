# faith
**faith** is a blockchain built using Cosmos SDK and Tendermint
We forked strange consumer chain.
These three remain: faith, hope and love. But the greatest of these is love.

## Get started

Install [go](https://go.dev/dl/)

## Build and install to go bin path

```
make install
```

## Initialize config

Come up with a moniker for your node, then run:

```
faithd init $MONIKER
```
 
 
 
## Launch with genesis file or run as standalone chain

To launch as a consumer chain, download and save shared genesis file to `~/.faith/config/genesis.json`. Additionally add peering information (`persistent_peers` or `seeds`) to `~/.faith/config/config.toml`

To instead launch as a standalone, single node chain, run:

```
faithd add-consumer-section
```

## Launch node

```
faithd start
```
