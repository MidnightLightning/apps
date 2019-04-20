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
    ensAddress = process.env.ENS,
    resolverAddress = process.env.RESOLVER,
    tld = process.env.TLD,
    rootName = process.env.ROOT_NAME,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  if(!ensAddress) log("missing ens address")
  if(!tld) log("missing tld")
  if(!rootName) log("missing rootName")

  const tldNode = namehash(tld)
  const label = '0x'+keccak256(rootName)
  const fullname = `${rootName}.${tld}`
  const node = namehash(fullname)

  const accounts = await getAccounts(web3)

  const ENS = artifacts.require('AbstractENS')

  const ens = await ENS.at(ensAddress)

  if(resolverAddress){
    resolverAddress = resolverAddress.toLowerCase()
  } else {
    resolverAddress = await ens.resolver(namehash('resolver.eth'))
  }
  log(resolverAddress)

  let currentResolver = await ens.resolver(node)

  if(currentResolver === resolverAddress) {
    log(`resolver for '${fullname}' is '${resolverAddress}'`)
  } else {
    log(`resolver is '${currentResolver}' and should be '${resolverAddress}'`)
    log(`setting resolver for '${fullname}' to '${resolverAddress}'`)
    if(await ens.owner(node) !== accounts[0]) {
      await ens.setSubnodeOwner(tldNode, label, accounts[0])
    }
    await ens.setResolver(node, resolverAddress)

    if(await ens.resolver(node) !== resolverAddress) {
      log("resolver not set")
      throw new Error("resolver not set")
    }
  }
}
