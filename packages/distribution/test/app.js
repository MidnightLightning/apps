const Distribution = artifacts.require('Distribution.sol')
const MiniMeToken = artifacts.require('MiniMeToken.sol')
const MiniMeTokenFactory = artifacts.require('MiniMeTokenFactory.sol')
const Registry = artifacts.require('Registry.sol')
const data = '0x1d2b477a59421ec4c27c69a72f9fcd71e970d55988425c63479ed846b50d7eb7	dumm02	5.4321E+22	["0xcef7bc68c4d8c3f5d62e3f8326cde2e15a8e3410bc1a8b5aaf69b2211c91663b","0x412ed3b771ffb3b248ae7888cb21e96e2eb8292aff60b8d6c16fa08d98b3889a","0x2a3341f2beb4fbcd626904d217abdf5951764e103c56632d02889fee3b1ba6f0","0x707c11d2fd3ffa1ee9d43effbe217a89492030ceb314813a1c8bab6b77effc3b","0x893406a968bb772f3eb544dcb3bd542bd791de0f3267934bfb44db7eebcedd65","0x995ac5cc14046860e3850ceb14df7185d8b520128f26c46bc2cf1ec5d5f47cbe","0x492d8d23793147559ce244c60e3e3f296652195b2fd5e7d5535499573597073c","0xc9fa5316da45a4558199927c1720fdcaa8998c7335939263a9778e07b0966456","0xbad79c8297bdd98bcbd5be1e966d0941e24e2af959be1d3d7ecfc55ba276ede6","0x777e98e1ebaa4750b5aef9522ac0e1dcc285acc66600b80396adae6e73eb64db"]'

const regRoot = "0xbbdacbe4195b9287f120e672eebc85fa1877ba9a0db7b1a3414d5846f08dc160"

contract('Distribution', (accounts) => {

  it('award', async () => {
      let claimArgs = data.split("\t")
      console.log(claimArgs)
      claimArgs[1] = web3.fromAscii(claimArgs[1])                // username
      claimArgs[2] = web3.toBigNumber(claimArgs[2]).toFixed()    // amount
      claimArgs[3] = JSON.parse(claimArgs[3])                    // merkle proof
      const distribution = await Distribution.deployed()
      let valid = await distribution.award(...claimArgs)
      console.log(valid)
  })
})
