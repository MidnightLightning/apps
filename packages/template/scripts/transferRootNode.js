const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle
const defaultOwner = process.env.OWNER
const defaultENSAddress = process.env.ENS
const defaultRegistryAddress = process.env.REGISTRY

const tld = namehash('eth')
const label = '0x'+keccak256('daonuts')
console.log("daonuts:", label)
const node = namehash('daonuts.eth')
console.log("daonuts.eth:", node)
console.log("resolver:", '0x'+keccak256('resolver'))
console.log("resolver.eth:", namehash('resolver.eth'))

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    web3 = globalWeb3,
    ensAddress = defaultENSAddress,
    registryAddress = defaultRegistryAddress,
    owner = defaultOwner,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  if(!owner) console.log("missing owner")
  if(!ensAddress) console.log("missing ens address")
  if(!registryAddress) console.log("missing registry address")

  const accounts = await getAccounts(web3)

  log(`Deploying daonuts.eth with ENS: ${ensAddress} and owner: ${owner}`)
  const ENS = artifacts.require('AbstractENS')
  const REGISTRY = artifacts.require('Registry')

  const ens = await ENS.at(ensAddress)
  console.log("after ens")
  const publicResolver = await ens.resolver(namehash('resolver.eth'))
  console.log("after publicResolver")
  const registry = await REGISTRY.at(registryAddress)
  console.log("after registry")

  log('assigning ENS name to registry')

  if (await ens.owner(node) === accounts[0]) {
    log('Transferring name ownership from deployer to registry')
    await ens.setOwner(node, registry.address)
  } else {
    log('Creating subdomain and assigning it to daonutsReg')
    try {
      await ens.setSubnodeOwner(tld, label, registry.address)
    } catch (err) {
      console.error(
        `Error: could not set the owner of 'aragonid.eth' on the given ENS instance`,
        `(${ensAddress}). Make sure you have ownership rights over the subdomain.`
      )
      throw err
    }
    let nodeOwner = await ens.owner(node)
    console.log(nodeOwner, registry.address)
    console.log(await ens.owner(node) === registry.address)
  }
}
