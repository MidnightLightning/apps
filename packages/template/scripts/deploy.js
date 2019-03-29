const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle
const defaultOwner = process.env.OWNER
const defaultENSAddress = process.env.ENS

const tld = namehash('eth')
const label = '0x'+keccak256('daonuts')
const node = namehash('daonuts.eth')

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    web3 = globalWeb3,
    ensAddress = defaultENSAddress,
    owner = defaultOwner,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  const accounts = await getAccounts(web3)

  log(`Deploying daonuts.eth with ENS: ${ensAddress} and owner: ${owner}`)
  const AppInstaller = artifacts.require('AppInstaller')
  const Template = artifacts.require('Template')

  const appInstaller = await appInstaller.new()
  await logDeploy(appInstaller, { verbose })
  const template = await Template.new(ensAddress)
  await logDeploy(template, { verbose })

  log('===========')
  log('Deployed AppInstaller:', appInstaller.address)
  log('Deployed Template:', template.address)
}
