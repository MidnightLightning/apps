const Web3 = require("web3")
const bases = require("bases")
const web3 = new Web3("wss://rinkeby.infura.io/ws")
const abi = require("../build/contracts/Tipping.json").abi
const Tipping = new web3.eth.Contract(abi, "0x418fD43FE85D2ee0e32789388a998e08a5f75E8B")

Tipping.events.Tip({fromBlock: 0}, (err, e)=>{
  if(err) throw err;
  // excuse the following line. e.returnValues should be used but isn't showing correct values. something i don't understand. so i'm just parsing raw
  let [from, to, amount, ctype, id] = e.raw.data.substring(2).match(/.{1,64}/g).map(h=>`0x${h}`)
  from = "0x"+from.slice(26,64)
  to = "0x"+to.slice(26,64)
  amount = web3.utils.toBN(amount)
  ctype = ["NONE", "COMMENT", "POST"][web3.utils.hexToNumber(ctype)]
  id = bases.toBase36(web3.utils.hexToNumber(id))
  console.log(`${from} tipped ${to} ${amount} for ${ctype}:${id}`)
  // > 0x7b6c819e9db25c302a9add821361bb95524023 tipped 0x6d609ee3e7ea328d0c52d046a68a47292fce1b 1000 for NONE:0
  // > 0x7b6c819e9db25c302a9add821361bb95524023 tipped 0x8401eb5ff34cc943f096a32ef3d5113febe8d4 2500 for COMMENT:ehrjx0r
})
