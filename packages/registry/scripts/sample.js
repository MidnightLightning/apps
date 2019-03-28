const Web3 = require("web3")
const web3 = new Web3("https://rinkeby.infura.io")
const abi = require("../build/contracts/Registry.json").abi
const Registry = new web3.eth.Contract(abi, "0x06c7FbBC86713457fEdC8605D3359fEdE72ED171")

// get registered address for u/dummy02
const usernameInHex = web3.utils.asciiToHex("dummy02")
Registry.methods.usernameToOwner(usernameInHex).call().then(console.log)
// > 0x7b6C819e9db25c302A9adD821361bB95524023D7

// results for u/non_registered_user
const nonregInHex = web3.utils.asciiToHex("non_registered_user")
Registry.methods.usernameToOwner(nonregInHex).call().then(console.log)
// > 0x0000000000000000000000000000000000000000
