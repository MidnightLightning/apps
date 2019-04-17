const namehash = require('eth-ens-namehash').hash

console.log("dummy01.daonuts.eth", namehash('dummy01.daonuts.eth'))
console.log("daonuts.eth", namehash('daonuts.eth'))
console.log("daonuts.test", namehash('daonuts.test'))
console.log("resolver.eth", namehash('resolver.eth'))
