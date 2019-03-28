const Web3 = require("web3")
const web3 = new Web3("http://localhost:8545")
const abi = require("../build/contracts/Hamburger.json").abi
const Hamburger = new web3.eth.Contract(abi, "0x0A614b4655a4130ece71D12d5939c1c46070010e")

let id = web3.utils.fromAscii("BANNER")
console.log(id)

Hamburger.methods.getAssetsCount().call()
  .then(res=>{
    console.log(res)
  })

// Hamburger.methods.assets(id).call()
//   .then(res=>{
//     console.log(res)
//   })
