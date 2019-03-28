import '@babel/polyfill'

import Aragon from '@aragon/client'

const app = new Aragon()

const initialState = {
  tips: []
}
app.store(async (state, event) => {
  if (state === null) state = initialState

  switch (event.event) {
    case 'Tip':
      if(!state.tips.find(t=>t.eventId===event.id)){
        let newTip = Object.assign(event.returnValues, {eventId:event.id})
        let tips = state.tips.slice(0)
        tips.unshift(newTip)
        return { tips }
      }
    default:
      return state
  }
})
