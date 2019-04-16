const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle
const defaultOwner = process.env.OWNER
const defaultENSAddress = process.env.ENS

const tld = 'test'
const tldNode = namehash(tld)
const name = 'daonuts'
const label = '0x'+keccak256(name)
console.log(`${name}:`, label)
const fullname = `${name}.${tld}`
const node = namehash(fullname)
console.log(`${fullname}:`, node)
console.log("resolver:", '0x'+keccak256('resolver'))
console.log("resolver.eth:", namehash('resolver.eth'))

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    web3 = globalWeb3,
    ensAddress = "0xe7410170f87102DF0055eB195163A03B7F2Bff4A",
    owner = "0x7b6C819e9db25c302A9adD821361bB95524023D7",
    verbose = true
  } = {}
) => {

  const accounts = await getAccounts(web3)
  const ENS = artifacts.require('AbstractENS')
  const FIFS = artifacts.require('FIFSRegistrar')
  const ens = await ENS.at(ensAddress)
  let tldOwner = await ens.owner(tldNode)
  console.log(tldOwner)
  const fifs = await FIFS.at(tldOwner)

  let nodeOwner = await ens.owner(node)

  if(nodeOwner !== owner) {
    console.log(`${owner} is owner of ${fullname}`)
  } else {
    console.log(`registering ${fullname}`)
    try {
      await fifs.register(label, accounts[0], {from: accounts[0]})
    } catch (e) {
      throw e
    }
  }
}
