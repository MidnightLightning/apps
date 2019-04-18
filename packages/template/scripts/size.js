let bytecode = require("../build/contracts/Template.json").deployedBytecode
let bytecode2 = require("../build/contracts/AppInstaller.json").deployedBytecode

console.log("Template", bytecode.length)
console.log("AppInstaller", bytecode.length)
