let bytecode = require("../build/contracts/TokenManager.json").deployedBytecode
let bytecode2 = require("../build/contracts/Token.json").deployedBytecode

console.log("TokenManager", bytecode.length)
console.log("Token", bytecode.length)
