import "@babel/polyfill"

import Aragon from "@aragon/client"

const app = new Aragon()

const initialState = {
  assetsCount: 0
}
app.store(async (state, event) => {
  if (state === null) state = initialState

  switch (event.event) {
    case "Transfer":
      // console.log(event)
      const assetsCount = await getAssetsCount()
      return { ...state, assetsCount, transfer: {id: event.returnValues._tokenId} }
    default:
      return state
  }
})
// Transfer
// Balance
// Price
// Data
// Tax

function getAssetsCount() {
  // Get current value from the contract by calling the public getter
  return new Promise(resolve => {
    app
      .call("assetsCount")
      .first()
      .map(value => parseInt(value, 10))
      .subscribe(resolve)
  })
}
