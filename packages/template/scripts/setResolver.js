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
  if(!tld) console.log("missing tld")
  if(!rootName) console.log("missing rootName")

  const tldNode = namehash(tld)
  const label = '0x'+keccak256(rootName)
  console.log(`${rootName}:`, label)
  const fullname = `${rootName}.${tld}`
  const node = namehash(fullname)
  console.log(`${fullname}:`, node)

  const accounts = await getAccounts(web3)

  const ENS = artifacts.require('AbstractENS')

  const ens = await ENS.at(ensAddress)

  if(!resolverAddress) {
    resolverAddress = await ens.resolver(namehash('resolver.eth'))
  }

  if(await ens.resolver(node) !== resolverAddress) {
    log(`setting resolver for '${fullname}' as: ${resolverAddress}`)
    if(await ens.owner(node) !== accounts[0]) {
      await ens.setSubnodeOwner(tldNode, label, accounts[0])
    }
    await ens.setResolver(node, resolverAddress)
  }

  if(ens.resolver(node) === resolverAddress) throw new Error("resolver not set")
}
