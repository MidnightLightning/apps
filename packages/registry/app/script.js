import '@babel/polyfill'

import Aragon from '@aragon/client'

const app = new Aragon()

const initialState = {
  rootsCount: 0,
  regChange: 0
}

let account

app.store(async (state, event) => {
  if (state === null) state = initialState

  switch (event.event) {
    case 'Registered':
      if(event.returnValues.owner === account){
        return { ...state, regChange: ++state.regChange }
      }
      break;
    case 'Deregistered':
      if(event.returnValues.owner === account){
        return { ...state, regChange: ++state.regChange }
      }
      break;
    case 'RegistrationPeriodStarted':
      const rootsCount = await getRootsCount()
      return { ...state, rootsCount }
    default:
      return state
  }
})

app.accounts().subscribe(accounts => {
  console.log("account changed", accounts[0])
  account = accounts[0]
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
