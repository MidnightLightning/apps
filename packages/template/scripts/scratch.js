const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle

const Promisify = (inner) =>
    new Promise((resolve, reject) =>
        inner((err, res) => {
            if (err) {
                reject(err);
            } else {
                resolve(res);
            }
        })
    );

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    appInstallerAddress="0x8a26D3Ca29440A631EdA90fd9f3Dd184994EA10e",
    web3 = globalWeb3,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  log(web3.version)

  const APP_INSTALLER = artifacts.require('AppInstaller')
  const appInstaller = await APP_INSTALLER.at(appInstallerAddress)
  log(appInstaller)
  let tokenEvents = await appInstaller.getPastEvents('CreatedToken', {fromBlock: 0, toBlock: 'latest'})
  // let tokenEventsFilter = appInstaller.CreatedToken({fromBlock: block.number})
  // let tokenEvents = await Promisify(cb => tokenEventsFilter.get(cb))
  log(tokenEvents)
}
