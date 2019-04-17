const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    web3 = globalWeb3,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  const tldNode = namehash("eth")
  const label = '0x'+keccak256("daonuts")
  const node = namehash("daonuts.eth")

  const accounts = await getAccounts(web3)
  const ENS = artifacts.require('AbstractENS')

  const ens = await ENS.at("0x5f6f7e8cc7346a11ca2def8f827b7a0b612c56a1")

  log('assigning ENS name to registry')

  await ens.setSubnodeOwner(tldNode, label, "0x3E28381C11bDd9921604b4fc8858cdBC44cf1C3C")
}
