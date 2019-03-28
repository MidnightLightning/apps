import React from 'react'
import {
  AragonApp,
  AppView,
  AppBar,
  BaseStyles,
  Main,
  Button,
  Text,
  Field,
  TextInput,
  Info,
  breakpoint,
  observe
} from '@aragon/ui'
import Aragon, { providers } from '@aragon/client'
import styled from 'styled-components'
import bases from 'bases'

const AppContainer = styled(AragonApp)`
  display: flex;
  align-items: center;
  justify-content: center;
`

export default class App extends React.Component {

  state = {username: '', balance: 0, recipient:'', owner: '', url:'', amount:0, panelOpen: false, panel: {title: "Peach"}}

  initDone = false

  componentDidUpdate = (prevProps) => {
    if(!this.initDone) {
      this.init()
    }
    if (this.props.userAccount !== prevProps.userAccount)
      this.getUsername(this.props.userAccount)
  }

  init = async () => {
    this.getUsername(this.props.userAccount)
    this.initDone = true
  }

  getUsername = async (account) => {
    console.log("getUsername")
    let username = await this.props.app.call('getUsername', account).toPromise()
    let balance = await this.props.app.call('balances', username).toPromise()
    balance = web3.toBigNumber(balance).div("1e+18").toNumber()
    if(username) this.setState({username: web3.toUtf8(username), balance})
  }

  handleUrlChange = (event) => {
    this.setState({url: event.target.value})
  }

  handleAmountChange = (event) => {
    this.setState({amount: event.target.value})
  }

  handleRecipientChange = async (event) => {
    let recipient = event.target.value
    this.setState({recipient})
    this.getRecipientAddress(recipient)
  }

  getRecipientAddress = async (recipient) => {
    let owner = await this.props.app.call('getOwner', web3.fromAscii(recipient)).toPromise()
    if(owner === "0x0000000000000000000000000000000000000000")
      owner = ""
    this.setState({owner})
  }

  handleMenuPanelOpen = () => {
    this.props.sendMessageToWrapper('menuPanel', true)
  }

  claim = () => {
    this.props.app.claim(this.props.userAccount)
  }

  submitTip = async (event) => {
    let cid=0, ctype=0;

    let recipient = web3.fromAscii(this.state.recipient);

    try {
      let parts = (new URL(this.state.url)).pathname.split("/").filter(a=>(!!a))
      if(parts.length === 6) {
        ctype = 1         // comment
        cid = bases.fromBase36(parts[5])
      } else if(parts.length === 5) {
        ctype = 2         // post
        cid = bases.fromBase36(parts[3])
      }
    } catch(e) {}


    let value = web3.toBigNumber(this.state.amount).mul("1e+18")

    console.log(cid, ctype, value)

    let tokenAddress = await this.props.app.call('currency').toPromise()

    console.log(tokenAddress)

    let intentParams = {
      token: { address: tokenAddress, value: value.toFixed() }
      // gas: 2000000
    }

    this.props.app.tip(recipient, value.toFixed(), ctype, cid.toString(), intentParams)
  }

  usePanel = (panel) => {
    this.setState({panelOpen: true, panel})
  }

  closePanel = () => {
    this.setState({panelOpen: false})
  }

  render () {
    let ownerText = this.state.owner
    if(this.props.userAccount && this.state.owner === this.props.userAccount)
      ownerText = "You"

    console.log("balance", typeof this.state.balance, this.state.balance, !!this.state.balance)
    return (
      <Main>
        <AppView appBar={<AppBar title="Tipping" />}>
          {!!this.state.username && <Welcome username={this.state.username} />}
          {!!this.state.balance && <Claim balance={this.state.balance} claim={this.claim} />}
          <Text size="xlarge">Tip for content (or not, it's also ok to leave url blank)</Text>
          <Field label="Recipient:">
            <TextInput placeholder="username" value={this.state.recipient} onChange={this.handleRecipientChange} />
            {!!this.state.recipient && <Info style={{"margin-top": "10px"}}>Owner: {ownerText || `${this.state.recipient} will need to register in order to claim this tip`}</Info>}
          </Field>
          <Field label="Amount:">
            <TextInput type="number" value={this.state.amount} onChange={this.handleAmountChange} />
          </Field>
          <Field label="Content Url:">
            <TextInput placeholder value={this.state.url} onChange={this.handleUrlChange} />
          </Field>
          <Button mode="strong" emphasis="positive" onClick={this.submitTip}>Tip</Button>
          <hr />
          <ObservedTips observable={this.props.observable} />
        </AppView>
      </Main>
    )
  }
}

const ObservedTips = observe(
  (state$) => state$,
  { tips: [] }
)(
  ({ tips }) => <TipList tips={tips} />
)

function TipList({tips}) {
  const types = ["NONE", "COMMENT", "POST"]
  const listItems = tips.map((tip) => {
    console.log(tip)
    return (
      <li>{`${web3.toUtf8(tip.fromName)} TIPPED ${web3.toUtf8(tip.toName)} ${web3.toBigNumber(tip.amount).div("1e+18").toFixed()} for ${types[tip.ctype]}:${bases.toBase36(tip.cid)}`}</li>
    )
  });
  return (
    <ul>{listItems}</ul>
  );
}

function Welcome({username}) {
  return (
    <div>
      <Text.Block style={{ textAlign: 'center' }} size='large'>welcome, </Text.Block>
      <Text.Block style={{ textAlign: 'center' }} size='xxlarge'>{username}</Text.Block>
    </div>
  )
}

function Claim({balance, claim}) {
  return (
    <Field label="Claim tips:">
      <Button mode="strong" emphasis="positive" onClick={claim}>Claim</Button>
      <Info.Action style={{"margin-top": "10px"}}>You have {balance} in tips to claim (you were tipped before registering)</Info.Action>
    </Field>
  )
}

// const Main = styled.div`
//   height: 100vh;
//   min-width: 320px;
// `

const Grid = styled.div`
  display: grid;
  grid-template-columns: 1fr;
  grid-auto-rows: 270px;
  grid-gap: 30px;

  ${breakpoint(
    'medium',
    `
      grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
     `,
  )};
`
