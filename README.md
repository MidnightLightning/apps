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
1. `npm run publish_apps`
1. `npm run template` (follow prompts)
1. dao interface should open in browser. alternatively read DAO_ADDRESS from `Created DAO: <DAO_ADDRESS>` and open `http://localhost:3000/#/<DAO_ADDRESS>` manually

## Develop with Docker

1. Run `docker-compose build`. This will create the `aragon-cli` and `aragon` Docker images (from `Dockerfile` and `Dockerfile.aragon` config files respectively)
1. Run `docker-compose up -d devchain`. This starts up the private Ethereum testnet instance. We need this up first as it creates some Ethereum test accounts and deploys an ENS instance upon startup.
1. Run `docker-compose logs devchain` and find the log line indicating "ENS instance deployed at <ENS_ADDRESS>". Copy the ENS address into the `.env` file for the `REACT_APP_ENS_REGISTRY_ADDRESS` value
1. Run `docker-compose up -d`. This will start up the remaining containers (`ipfs` and `aragon`).
1. Run `docker-compose logs -f` in a terminal to monitor the progress of all the containers.

### Install apps and dao from template

1. In another terminal, run `docker-compose run --rm aragon bash`. This will bring you into a bash shell, running inside the `aragon` container.
1. Run `cd /app && lerna link && lerna bootstrap --hoist` to get dependencies setup
1. Run `npm run publish_apps docker`.
1. Run `npm run template` and follow prompts

Read DAO_ADDRESS from `Created DAO: <DAO_ADDRESS>` and open `http://<DOCKER_HOST_IP>:3000/#/<DAO_ADDRESS>` manually

## Building Go packages

1. install solc v0.4.25
  1. `wget https://github.com/ethereum/solidity/releases/download/v0.4.25/solidity-ubuntu-trusty.zip`
  1. `unzip solidity-ubuntu-trusty`
  1. `sudo cp solidity-ubuntu-trusty/solc /usr/local/bin/`
