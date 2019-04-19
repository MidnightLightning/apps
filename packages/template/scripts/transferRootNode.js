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

  if(!ensAddress) console.log("missing ens address")
  if(!registryAddress) console.log("missing registry address")
  if(!tld) console.log("missing tld")
  if(!rootName) console.log("missing rootName")

  const tldNode = namehash(tld)
  const label = '0x'+keccak256(rootName)
  console.log(`${rootName}:`, label)
  const fullname = `${rootName}.${tld}`
  const node = namehash(fullname)
  console.log(`${fullname}:`, node)

  const accounts = await getAccounts(web3)

  log(`Deploying ${fullname} with ENS: ${ensAddress} and account: ${accounts[0]}`)
  const ENS = artifacts.require('AbstractENS')
  const REGISTRY = artifacts.require('Registry')

  const ens = await ENS.at(ensAddress)

  log('assigning ENS name to registry')

  if (await ens.owner(node) === accounts[0]) {
    log('Transferring name ownership from deployer to registry')
    await ens.setOwner(node, registryAddress)
  } else if (await ens.owner(tldNode) === accounts[0]) {
    log('Creating subdomain and assigning it to daonuts registry')
    await ens.setSubnodeOwner(tldNode, label, registryAddress)
  } else {
    throw new Error(`${accounts[0]} owns neither ${fullname} nor ${tld} to claim it. transfer ownership to this account first.`)
  }

  if(ens.resolver(node) === resolverAddress) throw new Error("resolver not set")
  if(ens.owner(node) !== registryAddress) throw new Error("registry not node owner")
}
