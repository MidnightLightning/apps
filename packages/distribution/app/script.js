import '@babel/polyfill'

import Aragon from '@aragon/client'

const app = new Aragon()

const initialState = {
  rootsCount: 0
}
app.store(async (state, event) => {
  if (state === null) state = initialState

  console.log(event)

  switch (event.event) {
    case 'DistributionStarted':
      const rootsCount = await getRootsCount()
      return { ...state, rootsCount }
    default:
      return state
  }
})

function getRootsCount() {
  // Get current value from the contract by calling the public getter
  return new Promise(resolve => {
    app
      .call('getRootsCount')
      .first()
      .map(value => parseInt(value, 10))
      .subscribe(resolve)
  })
}
