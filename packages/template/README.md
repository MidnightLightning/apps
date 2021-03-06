# Aragon React Boilerplate

> 🕵️ [Find more boilerplates using GitHub](https://github.com/search?q=topic:aragon-boilerplate) |
> ✨ [Official boilerplates](https://github.com/search?q=topic:aragon-boilerplate+org:aragon)

React boilerplate for Aragon applications.

This boilerplate includes a fully working example app, complete with a background worker and a front-end in React (with Aragon UI). Also comes with a DAO Template which will allow for using your app to interact with other Aragon apps like the Voting app. You can read more about DAO Template [here](https://hack.aragon.org/docs/templates-intro).

## Usage

To setup use the command `create-aragon-app`:

```sh
npx create-aragon-app <app-name> react
```

## Make the template work with your app

- In order for the template to work properly, it needs to know what the name of your app is. Replace `<app-name>` in [this line](https://github.com/aragon/aragon-react-boilerplate/blob/master/contracts/Template.sol#L68) with the name of your app in the `arapp.json` file (e.g. `app` for `app.aragonpm.eth`)

- Edit the roles defined in the template to configure your DAO as you want!

## Run the template

```sh
npx aragon run --template Template --template-init @ARAGON_ENS
```

## Running your app

### Using HTTP

Running your app using HTTP will allow for a faster development process of your app's front-end, as it can be hot-reloaded without the need to execute `aragon run` every time a change is made.

- First start your app's development server running `npm run start:app`, and keep that process running. By default it will rebuild the app and reload the server when changes to the source are made.

- After that, you can run `npm run start:http` or `npm run start:http:template` which will compile your app's contracts, publish the app locally and create a DAO. You will need to stop it and run it again after making changes to your smart contracts.

Changes to the app's background script (`app/script.js`) cannot be hot-reloaded, after making changes to the script, you will need to either restart the development server (`npm run start:app`) or rebuild the script `npm run build:script`.

### Using IPFS

Running your app using IPFS will mimic the production environment that will be used for running your app. `npm run start:ipfs` will run your app using IPFS. Whenever a change is made to any file in your front-end, a new version of the app needs to be published, so the command needs to be restarted.

## What's in this boilerplate?

### npm Scripts

- **start** or **start:ipfs**: Runs your app inside a DAO served from IPFS
- **start:http**: Runs your app inside a DAO served with HTTP (hot reloading)
- **start:ipfs:template**: Creates a DAO with the [Template](https://github.com/aragon/aragon-react-boilerplate/blob/master/contracts/Template.sol) and serves the app from IPFS
- **start:http:template**: Creates a DAO with the [Template](https://github.com/aragon/aragon-react-boilerplate/blob/master/contracts/Template.sol) and serves the app with HTTP (hot reloading)
- **prepare**: Install dependencies of the app
- **start:app**: Starts a development server for your app
- **compile**: Compile the smart contracts
- **build**: Builds the front-end and background script
- **test**: Runs tests for the contracts
- **publish:patch**: Release a patch version to aragonPM (only frontend/content changes allowed)
- **publish:minor**: Release a minor version to aragonPM (only frontend/content changes allowed)
- **publish:major**: Release a major version to aragonPM (frontend **and** contract changes)
- **versions**: Check the currently installed versions of the app

### Libraries

- [**@aragon/os**](https://github.com/aragon/aragonos): Aragon interfaces
- [**@aragon/api**](https://github.com/aragon/aragon.js/tree/master/packages/aragon-api): Wrapper for Aragon application RPC
- [**@aragon/ui**](https://github.com/aragon/aragon-ui): Aragon UI components (in React)

## What you can do with this boilerplate?

### Publish

You can publish you app on [aragonPM](https://hack.aragon.org/docs/apm). See how in our [publish guide](https://hack.aragon.org/docs/guides-publish).

> **Note**<br>
> The [Template](https://github.com/aragon/aragon-react-boilerplate/blob/master/contracts/Template.sol) will not be published.

### Using a different Ethereum account

You can use a different account to interact with you app. [Check the documentation](https://hack.aragon.org/docs/guides-faq#set-a-private-key).

### Propagate content

You can propagate the content of your app on IPFS. Learn more in our [troubleshooting guide](https://hack.aragon.org/docs/guides-faq#propagating-your-content-hash-through-ipfs) or use the `aragon ipfs propagate` command:

```
npx aragon ipfs propagate <cid>
```

Where `cid` is your content id hash (this will be displayed after publishing).
