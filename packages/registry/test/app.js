const Registry = artifacts.require('Registry.sol')

contract('Registry', (accounts) => {
  // it('should be tested')

  it('hash', async () => {
      const username = "dummy01"
      const registry = await Registry.new()
      let hash = await registry.hash("0xb4124cEB3451635DAcedd11767f004d8a28c6eE7", username)
      console.log("hash is:", hash)
      let addressHash = await registry.hashAddress("0xb4124cEB3451635DAcedd11767f004d8a28c6eE7")
      console.log("addressHash is:", addressHash)
      let usernameHash = await registry.hashBytes32(username)
      console.log("usernameHash is:", usernameHash)
  })
})
