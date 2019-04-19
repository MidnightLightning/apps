const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle
const defaultENSAddress = process.env.ENS
const defaultRegistryAddress = process.env.REGISTRY
const defaultResolverAddress = process.env.RESOLVER
const defaultTLD = process.env.TLD
const defaultRootName = process.env.ROOT_NAME

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    web3 = globalWeb3,
    ensAddress = defaultENSAddress,
    registryAddress = defaultRegistryAddress,
    resolverAddress = defaultResolverAddress,
    tld = defaultTLD,
    rootName = defaultRootName,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  if(!ensAddress) log("missing ens address")
  if(!registryAddress) log("missing registry address")
  if(!tld) log("missing tld")
  if(!rootName) log("missing rootName")

  const tldNode = namehash(tld)
  const label = '0x'+keccak256(rootName)
  const fullname = `${rootName}.${tld}`
  const node = namehash(fullname)

  const accounts = await getAccounts(web3)
  const ENS = artifacts.require('AbstractENS')
  const REGISTRY = artifacts.require('Registry')
  const ens = await ENS.at(ensAddress)

  const resolver = await ens.resolver(node)

  log(resolver)
  if(!resolver) {
    log(`resolver not set`)
    throw new Error("resolver not set")
  }

  log('assigning ENS name to registry')

  if (await ens.owner(node) === accounts[0]) {
    log(`Transferring ${fullname} ownership from deployer to registry:${registryAddress}`)
    await ens.setOwner(node, registryAddress)
  } else if (await ens.owner(tldNode) === accounts[0]) {
    log(`Creating ${fullname} and assigning it to registry:${registryAddress}`)
    await ens.setSubnodeOwner(tldNode, label, registryAddress)
  } else {
    log(`${accounts[0]} owns neither ${fullname} nor ${tld} to claim it. cannot transfer ownership.`)
    throw new Error(`${accounts[0]} owns neither ${fullname} nor ${tld} to claim it. cannot transfer ownership.`)
  }

  if(await ens.owner(node) === registryAddress) {
    log(`registry is node owner`)
  } else {
    log(`registry is not node owner`)
    throw new Error("registry is not node owner")
  }
}
