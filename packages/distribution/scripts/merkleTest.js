const utils = require("ethereumjs-util")
const setLengthLeft = utils.setLengthLeft
const setLengthRight = utils.setLengthRight
let addressBuffer = utils.toBuffer("0xb4124cEB3451635DAcedd11767f004d8a28c6eE7")
let uintBuffer = setLengthLeft(utils.toBuffer(340), 32)
let hashBuffer = Buffer.concat([addressBuffer,uintBuffer])
console.log("addressHash", utils.bufferToHex(utils.keccak256(addressBuffer)))
console.log("uintHashSLL", utils.bufferToHex(utils.keccak256(uintBuffer)))
console.log("hash", utils.bufferToHex(utils.keccak256(hashBuffer)))
