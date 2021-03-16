# Doru

Doru provides orchestration for a decentralised applications to compute and
store data at the edge. Edge devices are both browsers for interfacing with people,
and any device at the edge that can measure, compute or perform.

## Doru wallet

Doru SDK provides user-owned identities for web3 dApps for private, collaborative
and collective problems, building on
 - textile/threads for data orchestration,
 - libp2p/ipfs for networking/data traversal
 - ethereum for global synchronisation
 - OpenST for payment rails
 - fission/ucan for permission delegation

Doru provides a 6digit PIN fully-encrypted no-signup account management
and recovery flow in the browser for users.

Stateful DApps can be supported by any number of serverless providers
the user delegates permission to. The user pays providers over payment channels
for (encrypted) data usage by the providers. This provides a revenue stream for developers
who can host default providers for their dApp.

Doru SDK allows a web wallet for interacting with Ethereum (layer 1 and 2)

## Getting started
### Running development

Run Doru and Threads with docker-compose. A Makefile has the default commands.

Create an environment `.env` file with following variables set (see `.env.example`)

```
REPO_PATH=~/mydatarepo
```
and run
```
make doru-up
```
