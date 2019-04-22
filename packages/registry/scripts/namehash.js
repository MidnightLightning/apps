const namehash = require('eth-ens-namehash').hash

console.log("dummy01.daonuts.eth", namehash('dummy01.daonuts.eth'))
console.log("daonuts.eth", namehash('daonuts.eth'))
console.log("daonuts.test", namehash('daonuts.test'))
console.log("resolver.eth", namehash('resolver.eth'))
console.log("aragonpm.eth", namehash('aragonpm.eth'))
console.log("open.aragonpm.eth", namehash('open.aragonpm.eth'))
console.log("daonuts-token-manager.aragonpm.eth", namehash('daonuts-token-manager.aragonpm.eth'))
console.log("daonuts-token-manager.open.aragonpm.eth", namehash('daonuts-token-manager.open.aragonpm.eth'))
