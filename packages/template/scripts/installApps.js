const namehash = require('eth-ens-namehash').hash
const keccak256 = require('js-sha3').keccak_256
const logDeploy = require('@aragon/os/scripts/helpers/deploy-logger')
const getAccounts = require('@aragon/os/scripts/helpers/get-accounts')

const globalArtifacts = this.artifacts // Not injected unless called directly via truffle
const globalWeb3 = this.web3 // Not injected unless called directly via truffle

const Promisify = (inner) =>
    new Promise((resolve, reject) =>
        inner((err, res) => {
            if (err) {
                reject(err);
            } else {
                resolve(res);
            }
        })
    );

module.exports = async (
  truffleExecCallback,
  {
    artifacts = globalArtifacts,
    web3 = globalWeb3,
    daoAddress = process.env.DAO,
    ensAddress = process.env.ENS,
    appInstallerAddress = process.env.APP_INSTALLER,
    permissionSetterAddress = process.env.PERMISSION_SETTER,
    regRoot = process.env.REG_ROOT,
    distRoot = process.env.DIST_ROOT,
    rootNode = process.env.ROOT_NODE,
    rootName = process.env.ROOT_NAME,
    tld = process.env.TLD,
    apmRootName = process.env.APM_ROOT_NAME,
    verbose = true
  } = {}
) => {
  const log = (...args) => {
    if (verbose) { console.log(...args) }
  }

  if(!daoAddress) log("missing app dao address")
  if(!appInstallerAddress) log("missing appInstaller address")
  if(!permissionSetterAddress) log("missing permissionSetter address")

  log(web3.version)

  let block = await web3.eth.getBlock("latest")

  const APP_INSTALLER = artifacts.require('AppInstaller')
  const PERMISSION_SETTER = artifacts.require('PermissionSetter')
  const TOKEN = artifacts.require('Token')

  const accounts = await getAccounts(web3)
  const appInstaller = await APP_INSTALLER.at(appInstallerAddress)
  const permissionSetter = await PERMISSION_SETTER.at(permissionSetterAddress)

  let tokenManagerAppId = namehash(`daonuts-token-manager.${apmRootName}`)

  // installCurrencyManager
  let currencyName = "Currency"
  try {
    let gas = await appInstaller.installCurrencyManager.estimateGas(daoAddress, tokenManagerAppId, currencyName, "NUTS")
    log(`'installCurrencyManager' gas:`, gas)
    await appInstaller.installCurrencyManager(daoAddress, tokenManagerAppId, currencyName, "NUTS")
  } catch(e){
    log(e)
  }

  // installKarmaManager
  let karmaName = "Karma"
  try {
    let gas = await appInstaller.installKarmaManager.estimateGas(daoAddress, tokenManagerAppId, karmaName, "KARMA")
    log(`'installKarmaManager' gas:`, gas)
    await appInstaller.installKarmaManager(daoAddress, tokenManagerAppId, karmaName, "KARMA")
  } catch(e){
    log(e)
  }

  let currencyAddress, karmaAddress
  try {
    let tokenEvents = await appInstaller.getPastEvents('CreatedToken', {fromBlock: block.number, toBlock: 'latest'})
    currencyAddress = tokenEvents.find(e=>e.returnValues.name===currencyName).returnValues.token
    karmaAddress = tokenEvents.find(e=>e.returnValues.name===karmaName).returnValues.token
  } catch(e){
    log(e)
  }

  log("tokens", currencyAddress, karmaAddress)

  let currency = await TOKEN.at(currencyAddress)
  let karma = await TOKEN.at(karmaAddress)

  // TODO - alt. get these from the token controller values
  let currencyManagerAddress = await currency.controller()
  let karmaManagerAddress = await karma.controller()

  // installVoting
  let votingAppId = namehash(`daonuts-karma-cap-voting.${apmRootName}`)
  try {
    let gas = await appInstaller.installVoting.estimateGas(daoAddress, votingAppId, currencyAddress, karmaAddress)
    log(`'installVoting' gas:`, gas)
    await appInstaller.installVoting(daoAddress, votingAppId, currencyAddress, karmaAddress)
  } catch(e){
    log(e)
  }

  // installTipping
  let tippingAppId = namehash(`daonuts-tipping.${apmRootName}`)
  try {
    let gas = await appInstaller.installTipping.estimateGas(daoAddress, tippingAppId, currencyAddress)
    log(`'installTipping' gas:`, gas)
    await appInstaller.installTipping(daoAddress, tippingAppId, currencyAddress)
  } catch(e){
    log(e)
  }

  // installRegistry
  let registryAppId = namehash(`daonuts-registry.${apmRootName}`)
  try {
    let gas = await appInstaller.installRegistry.estimateGas(daoAddress, registryAppId, rootNode, regRoot)
    log(`'installRegistry' gas:`, gas)
    await appInstaller.installRegistry(daoAddress, registryAppId, rootNode, regRoot)
  } catch(e){
    log(e)
  }

  // installDistribution
  let distributionAppId = namehash(`daonuts-distribution.${apmRootName}`)
  try {
    let gas = await appInstaller.installDistribution.estimateGas(daoAddress, distributionAppId, currencyManagerAddress, karmaManagerAddress, distRoot)
    log(`'installDistribution' gas:`, gas)
    await appInstaller.installDistribution(daoAddress, distributionAppId, currencyManagerAddress, karmaManagerAddress, distRoot)
  } catch(e){
    log(e)
  }

  // installHamburger
  let hamburgerAppId = namehash(`daonuts-hamburger.${apmRootName}`)
  try {
    let gas = await appInstaller.installHamburger.estimateGas(daoAddress, hamburgerAppId, currencyManagerAddress)
    log(`'installHamburger' gas:`, gas)
    await appInstaller.installHamburger(daoAddress, hamburgerAppId, currencyManagerAddress)
  } catch(e){
    log(e)
  }

  let votingAddress, tippingAddress, registryAddress, distributionAddress, hamburgerAddress
  try {
    let appEvents = await appInstaller.getPastEvents('InstalledApp', {fromBlock: block.number, toBlock: 'latest'})
    votingAddress = appEvents.find(e=>e.returnValues.appId===votingAppId).returnValues.appProxy
    tippingAddress = appEvents.find(e=>e.returnValues.appId===tippingAppId).returnValues.appProxy
    registryAddress = appEvents.find(e=>e.returnValues.appId===registryAppId).returnValues.appProxy
    distributionAddress = appEvents.find(e=>e.returnValues.appId===distributionAppId).returnValues.appProxy
    hamburgerAddress = appEvents.find(e=>e.returnValues.appId===hamburgerAppId).returnValues.appProxy
  } catch(e){
    log(e)
  }

  // setPermissions
  try {
    let gas = await permissionSetter.setPermissions.estimateGas(daoAddress, currencyManagerAddress, karmaManagerAddress, votingAddress, registryAddress, distributionAddress, hamburgerAddress, tippingAddress)
    log(`'setPermissions' gas:`, gas)
    await permissionSetter.setPermissions(daoAddress, currencyManagerAddress, karmaManagerAddress, votingAddress, registryAddress, distributionAddress, hamburgerAddress, tippingAddress)
  } catch(e){
    log(e)
  }

  if(!ensAddress) log("missing ens address")
  if(!registryAddress) log("missing registry address")
  if(!rootName) log("missing rootName")
  if(!tld) log("missing tld")

  const tldNode = namehash(tld)
  const label = '0x'+keccak256(rootName)
  const fullname = `${rootName}.${tld}`
  const node = namehash(fullname)

  const ENS = artifacts.require('AbstractENS')
  const ens = await ENS.at(ensAddress)

  const resolver = await ens.resolver(node)

  if(!resolver) {
    log(`resolver not set`)
    throw new Error("resolver not set")
  }

  log('assigning ENS name to registry')

  log(ensAddress, resolver, registryAddress, await ens.owner(node), accounts[0])

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

  log("finished installer")
}
