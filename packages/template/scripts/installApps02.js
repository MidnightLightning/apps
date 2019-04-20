const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle
const defaultDAOAddress = process.env.DAO
const defaultRegRoot = process.env.REG_ROOT
const defaultDistRoot = process.env.DIST_ROOT
const defaultAppInstaller02Address = process.env.APP_INSTALLER_02
const defaultCurrencyManagerAddress = process.env.CURRENCY_MANAGER
const defaultKarmaManagerAddress = process.env.KARMA_MANAGER
const defaultVotingAddress = process.env.VOTING

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    web3 = globalWeb3,
    daoAddress = defaultDAOAddress,
    regRoot = defaultRegRoot,
    distRoot = defaultDistRoot,
    appInstaller02Address = defaultAppInstaller02Address,
    currencyManagerAddress = defaultCurrencyManagerAddress,
    karmaManagerAddress = defaultKarmaManagerAddress,
    votingAddress = defaultVotingAddress,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  if(!appInstaller02Address) log("missing app installer02 address")
  if(!currencyManagerAddress) log("missing app currencyManager address")
  if(!karmaManagerAddress) log("missing app karmaManager address")
  if(!votingAddress) log("missing app voting address")
  const APP_INSTALLER_02 = artifacts.require('AppInstaller02')

  const appInstaller02 = await APP_INSTALLER_02.at(appInstaller02Address)

  // installer02 apps
  try {
    let gas = await appInstaller02.install.estimateGas(daoAddress, currencyManagerAddress, karmaManagerAddress, votingAddress, regRoot, distRoot)
    log(`'installer02' gas:`, gas)
    await appInstaller02.install(daoAddress, currencyManagerAddress, karmaManagerAddress, votingAddress, regRoot, distRoot)
  } catch(e){
    log(e)
  }

    log("finished installer02")
}
