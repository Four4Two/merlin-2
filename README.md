# Merlin

![Banner!](assets/banner.png)

[![Project Status: Active -- The project has reached a stable, usable
state and is being actively
developed.](https://img.shields.io/badge/repo%20status-Active-green.svg?style=flat-square)](https://www.repostatus.org/#active)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/four4two/merlin/v11)
[![Go Report
Card](https://goreportcard.com/badge/github.com/four4two/merlin?style=flat-square)](https://goreportcard.com/report/github.com/four4two/merlin/v11)
[![Version](https://img.shields.io/github/tag/four4two/merlin.svg?style=flat-square)](https://github.com/four4two/merlin/releases/latest)
[![License:
Apache-2.0](https://img.shields.io/github/license/four4two/merlin.svg?style=flat-square)](https://github.com/four4two/merlin/blob/main/LICENSE)
[![Lines Of
Code](https://img.shields.io/tokei/lines/github/four4two/merlin?style=flat-square)](https://github.com/four4two/merlin)
[![GitHub
Super-Linter](https://img.shields.io/github/actions/workflow/status/four4two/merlin/lint.yml?style=flat-square&label=Lint)](https://github.com/marketplace/actions/super-linter)
[![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)](https://discord.gg/merlin)

Merlin is a fair-launched, customizable automated market maker for
interchain assets that allows the creation and management of
non-custodial, self-balancing, interchain token index similar to one of
Balancer.

Inspired by [Balancer](http://balancer.finance/whitepaper) and Sunny
Aggarwal's '[DAOfying Uniswap Automated Market Maker
Pools](https://www.sunnya97.com/blog/daoifying-uniswap-automated-market-maker-pools)',
the goal for Merlin is to provide the best-in-class tools that extend
the use of AMMs within the Cosmos ecosystem beyond traditional token
swap-type use cases. Bonding curves, while have found its primary use
case in decentralized exchange mechanisms, its potential use case can be
further extended through the customizability that Merlin offers.
Through the customizability offered by Merlin such as custom-curve AMMs,
dynamic adjustments of spread factors, multi-token liquidity pools--the AMM
can offer decentralized formation of token fundraisers, interchain
staking, options market, and more for the Cosmos ecosystem.

Whereas most Cosmos zones have focused their incentive scheme on the
delegators, Merlin attempts to align the interests of multiple
stakeholders of the ecosystem such as LPs, DAO members, as well as
delegators. One mechanism that is introduced is how staked liquidity
providers have sovereign ownership over their pools, and through the
pool governance process allow them to adjust the parameters depending on
the pool's competition and market conditions. Merlin is a sovereign
Cosmos zone that derives its sovereignty not only from its
application-specific blockchain architecture but also the collective
sovereignty of the LPs that has aligned interest to different tokens
that they are providing liquidity for.

## System Requirements

This system spec has been tested by many users and validators and found
to be comfortable:

- Quad Core or larger AMD or Intel (amd64) CPU
  - ARM CPUs like the Apple M1 are not supported at this time.
- 64GB RAM (A lot can be in swap)
- 1TB NVMe Storage
- 100MBPS bidirectional internet connection

You can run Merlin on lower-spec hardware for each component, but you
may find that it is not highly performant or prone to crashing.

## Documentation

For the most up to date documentation please visit
[docs.merlin.zone](https://docs.merlin.zone/)

## Joining the Mainnet

[Please visit the official instructions on how to join the Mainnet
here.](https://docs.merlin.zone/networks/join-mainnet)

Thank you for supporting a healthy blockchain network and community by
running an Merlin node!

## Contributing

The contributing guide for Merlin explains the branching structure, how
to use the SDK fork, and how to make / test updates to SDK branches.

## LocalMerlin

LocalMerlin is a containerized local Merlin testnet used for trying out new features locally. 
LocalMerlin documentation can be found [here](https://github.com/four4two/merlin/tree/main/tests/localmerlin)
