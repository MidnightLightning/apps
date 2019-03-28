# daonuts Aragon kit

### requirements
1. `npm i -g @aragon/cli`
1. `aragon devchain` in one terminal
1. `aragon ipfs` in another terminal
1. clone [daonuts-aragon-apps](https://github.com/daonuts/daonuts-aragon-apps)
1. `cd aragon-apps/apps/voting` && `npm i` && `npm run compile` && `npm run publish:patch`
1. clone [distribution](https://github.com/daonuts/distribution)
1. `cd distribution` && `npm i` && `npm run publish:patch`
1. clone [tipping](https://github.com/daonuts/tipping)
1. `cd tipping` && `npm i` && `npm run publish:patch`
1. clone [registry](https://github.com/daonuts/registry)
1. `cd registry` && `npm i` && `npm run publish:patch`

### publish kit
1. clone this repo && `npm i`
1. `aragon deploy Kit --init @ARAGON_ENS`
1. `aragon apm publish [patch|minor|major] [kit-address]`

### create dao from kit
| WARNING: using the kit is currently broken - it takes too much gas. |
| --- |
1. `dao new --kit daonuts-kit --fn newInstance` # get dao-address from output

### create dao with "interactive manual deploy"
1. `cd [this repo directory]`
1. `. ./scripts/interactive_manual_deploy.sh 1 rinkeby`

### view dao

1. clone [aragon repo](https://github.com/aragon/aragon) && `npm i`
1. in aragon directory: `npm run start:local`
1. open http://localhost:3000/#/[dao-address]
