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
    owner = defaultOwner,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  const accounts = await getAccounts(web3)

  log(`Deploying daonuts.eth with ENS: ${ensAddress} and owner: ${owner}`)
  const Registry = artifacts.require('Registry')
  const ENS = artifacts.require('AbstractENS')

  const ens = await ENS.at(ensAddress)
  const publicResolver = await ens.resolver(namehash('resolver.eth'))
  const daonutsReg = await Registry.new()
  await logDeploy(daonutsReg, { verbose })

  log('assigning ENS name to daonutsReg')

  if (await ens.owner(node) === accounts[0]) {
    log('Transferring name ownership from deployer to daonutsReg')
    await ens.setOwner(node, daonutsReg.address)
  } else {
    log('Creating subdomain and assigning it to daonutsReg')
    try {
      await ens.setSubnodeOwner(tld, label, daonutsReg.address)
    } catch (err) {
      console.error(
        `Error: could not set the owner of 'aragonid.eth' on the given ENS instance`,
        `(${ensAddress}). Make sure you have ownership rights over the subdomain.`
      )
      throw err
    }
    console.log(await ens.owner(node), daonutsReg.address)
    console.log(await ens.owner(node) === daonutsReg.address)
  }

  // if (owner) {
  //   log('assigning owner name')
  //   await daonutsReg.register('0x'+keccak256('owner'), owner)
  // }

  log('===========')
  log('Deployed Registry:', daonutsReg.address)
}
