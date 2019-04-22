# daonuts

## Introduction

This is a monorepo for modules and apps related to the [/r/daonuts](https://www.reddit.com/r/daonuts) project.

## Deploy locally

### Install and run prerequisites

1. `npm i -g @aragon/cli`
1. `aragon devchain`  (in own terminal)
1. `aragon ipfs`  (in own terminal)
1. `git clone git@github.com:aragon/aragon.git && cd aragon`
1. `npm i && npm run start:local`

### Install apps and dao from template

1. `npm i -g lerna`
1. `git clone git@github.com:daonuts/apps.git && cd apps`
1. `lerna link`
1. `lerna bootstrap --hoist`
1. `cd packages/template`
1. `./scripts/publish_apps.sh`
1. `./scripts/create_dao.sh`
1. dao interface should open in browser. alternatively read DAO_ADDRESS from `Created DAO: <DAO_ADDRESS>` and open `http://localhost:3000/#/<DAO_ADDRESS>` manually

## Building Go packages

1. install solc v0.4.25
  1. `wget https://github.com/ethereum/solidity/releases/download/v0.4.25/solidity-ubuntu-trusty.zip`
  1. `unzip solidity-ubuntu-trusty`
  1. `sudo cp solidity-ubuntu-trusty/solc /usr/local/bin/`
