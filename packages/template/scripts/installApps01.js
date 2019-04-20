const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle
const defaultDAOAddress = process.env.DAO
const defaultAppInstaller01Address = process.env.APP_INSTALLER_01
const defaultAppInstaller02Address = process.env.APP_INSTALLER_02

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    web3 = globalWeb3,
    daoAddress = defaultDAOAddress,
    appInstaller01Address = defaultAppInstaller01Address,
    appInstaller02Address = defaultAppInstaller02Address,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  if(!daoAddress) log("missing app dao address")
  if(!appInstaller01Address) log("missing app installer01 address")
  if(!appInstaller02Address) log("missing app installer02 address")
  const APP_INSTALLER_01 = artifacts.require('AppInstaller01')

  const appInstaller01 = await APP_INSTALLER_01.at(appInstaller01Address)

  try {
    let gas = await appInstaller01.createTokens.estimateGas()
    log(`'createTokens' gas:`, gas)
    await await appInstaller01.createTokens()
  } catch(e){
    log(e)
  }

  const currency = await appInstaller01.currency()
  log(`currency created at ${currency}`)
  const karma = await appInstaller01.karma()
  log(`karma created at ${karma}`)

  // installer01 apps
  try {
    let gas = await appInstaller01.install.estimateGas(daoAddress, appInstaller02Address)
    log(`'installer01' gas:`, gas)
    await appInstaller01.install(daoAddress, appInstaller02Address)
  } catch(e){
    log(e)
  }

  log("finished installer01")
}
