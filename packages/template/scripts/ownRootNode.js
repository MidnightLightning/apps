const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle
const defaultENSAddress = process.env.ENS
const defaultTLD = process.env.TLD
const defaultRootName = process.env.ROOT_NAME

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    web3 = globalWeb3,
    ensAddress = defaultENSAddress,
    tld = defaultTLD,
    rootName = defaultRootName,
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

  let nodeOwner = await ens.owner(node)
  log(`current owner of '${fullname}': ${nodeOwner}`)

  if(nodeOwner !== accounts[0]) {
    if (await ens.owner(tldNode) === accounts[0]) {
      log(`${accounts[0]} owns '.${tld}'. claiming ${fullname}`)
      await ens.setSubnodeOwner(tldNode, label, accounts[0])
    } else if(!nodeOwner && process.env.NETWORK === 'rinkeby'){
      log(`registering '${fullname}'`)
      const FIFS = artifacts.require('FIFSRegistrar')
      let tldOwner = await ens.owner(tldNode)
      const fifs = await FIFS.at(tldOwner)
      try {
        await fifs.register(label, accounts[0], {from: accounts[0]})
      } catch (e) {
        log(e)
        throw e
      }
    }
  }

  if(await ens.owner(node) === accounts[0]) {
    log(`${accounts[0]} is owner of '${fullname}'`)
  } else {
    log(`${accounts[0]} is not owner of '${fullname}'`)
    throw new Error(`deployer is not node owner`)
  }
}
