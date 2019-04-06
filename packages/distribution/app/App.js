import React from 'react'
import {
  AragonApp,
  AppBar,
  AppView,
  SidePanel,
  Button,
  BaseStyles,
  Main,
  Text,
  Field,
  TextInput,
  Info,
  theme,
  breakpoint,
  observe
} from '@aragon/ui'
import Aragon, { providers } from '@aragon/client'
import styled from 'styled-components'

const AppContainer = styled(AragonApp)`
  display: flex;
  align-items: center;
  justify-content: center;
`

import dist01 from '../distributions/post/0x3c29d249.json'

export default class App extends React.Component {

  state = {username: '', roots: [], claimed: [], registry: null, panelOpen: false, panel: {title: "Peach"}}

  lastObservable = {}

  distributions = {
    "0x3c29d249": dist01
  }

  initDone = false

  componentDidUpdate = (prevProps) => {
    if(!this.initDone) {
      this.init()
    }
    if (this.props.userAccount !== prevProps.userAccount)
      this.getUsername(this.props.userAccount)
    if (this.props.observable && !prevProps.observable) {
      this.props.observable.subscribe(o=>{
        if(!o) return
        if(o.rootsCount !== this.lastObservable.rootsCount)
          this.getRoots(o.rootsCount)
        this.lastObservable = o
      })
    }
  }

  init = async () => {
    this.getRegistry()
    this.getUsername(this.props.userAccount)
    let rootsCount = await this.props.app.call('getRootsCount').toPromise()
    this.getRoots(rootsCount)
    this.initDone = true
  }

  handleMenuPanelOpen = () => {
    this.props.sendMessageToWrapper('menuPanel', true)
  }

  getRoots = async (length) => {
    let roots = await Promise.all(
        Array
          .from({length})
          .map((v,i)=>this.props.app.call('roots', i).toPromise())
      )

    this.setState({roots})
    this.getClaimed()
  }

  getClaimed = async () => {
    let username = this.state.username
    let roots = this.state.roots
    if(!username || !roots.length) return;
    let claimed = await Promise.all(
        roots.map((root,i)=>this.props.app.call('claimed', root, username).toPromise())
      )
    this.setState({claimed})
  }

  getUsername = async (account) => {
    console.log("getUsername")
    let username = await this.props.app.call('nameOfOwner', account).toPromise()
    if(username) this.setState({username})
    this.getClaimed()
  }

  getRegistry = async () => {
    console.log('getRegistry')
    let registry = await this.props.app.call('registry').toPromise()
    console.log(registry)
    this.setState({registry})
  }

  usePanel = (panel) => {
    this.setState({panelOpen: true, panel})
  }

  closePanel = () => {
    this.setState({panelOpen: false})
  }

  claim = (root, userDistData) => {
    let award = web3.toBigNumber(userDistData.award).toFixed()    // amount
    this.props.app.award(root, userDistData.username, award, userDistData.proof)
  }

  openClaimControls = (root) => {
    // console.log(this, root)
    this.usePanel({
      title: `Claim in distribution: ${root.slice(0,7)}...`,
      child: Claim,
      childProps: {
        app: this.props.app,
        root
      }
    })
  }

  openNewDistributionControls = () => {
    this.usePanel({
      title:"New asset",
      child: NewDistribution,
      childProps: {
        app: this.props.app
      }
    })
  }

  render () {
    const Child = this.state.panel.child
    return (
      <Main>
        <AppView appBar={<AppBar title="Distribution" endContent={<Button mode="strong" onClick={this.openNewDistributionControls}>New Distribution</Button>} />} >
          {this.state.username ? <Welcome username={this.state.username} /> : null}
          <Text size="xlarge">Accepted merkle roots:</Text>
          <RootList roots={this.state.roots} claim={this.claim} claimed={this.state.claimed} distributions={this.distributions} userAccount={this.props.userAccount} username={this.state.username} />
          <Info style={{"margin-top": "10px"}}>registry: {this.state.registry}</Info>
          <SidePanel title={this.state.panel.title} opened={this.state.panelOpen} onClose={this.closePanel}>
            {Child && <Child {...this.state.panel.childProps} />}
          </SidePanel>
        </AppView>
      </Main>
    )
  }
}

class NewDistribution extends React.Component {

  state = {newRoot: ''}

  handleNewRootChange = (event) => {
    this.setState({newRoot: event.target.value});
  }

  submitRoot = (event) => {
    let peach = this.props.app.addRoot(this.state.newRoot);
  }

  render() {
    return (
      <Field label="New distribution merkle root:">
        <TextInput value={this.state.newRoot} onChange={this.handleNewRootChange} />
        <Button onClick={this.submitRoot}>Add</Button>
      </Field>
    )
  }
}

function RootList({roots, scope, claim, claimed, distributions, userAccount, username}) {
  const listItems = roots.map((root, i) => {
    let dist = distributions[root.slice(0,10)]
    let userDistData
    if(userAccount)
      userDistData = dist.find(d=>d.username===username)
    return (
      <Card>
        <Content>
          <Label>
            <Text color={theme.textTertiary}>#{i} </Text>
            <span>{`${root.slice(0,10)}...`}</span>
          </Label>
          <Button disabled={!userDistData || !username || claimed[i]} mode="strong" emphasis="positive" onClick={()=>claim(root, userDistData)}>Claim</Button>
          {username && userDistData && !claimed[i] && <Info.Action style={{"margin-top": "10px"}}>You can claim {web3.toBigNumber(userDistData.award).div("1e+18").toFixed()}</Info.Action>}
          {username && !userDistData && <Info.Alert style={{"margin-top": "10px"}}>{username} not found in distribution</Info.Alert>}
          {userAccount && !username && <Info.Alert style={{"margin-top": "10px"}}>{userAccount.slice(0,8)}... has not registered</Info.Alert>}
          {claimed[i] && <Info style={{"margin-top": "10px"}}>Congrats, you claimed for this distribution</Info>}
        </Content>
      </Card>
    )
  });
  return (
    <Grid>{listItems}</Grid>
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

const Card = styled.div`
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 20px 30px;
  background: #ffffff;
  border: 1px solid rgba(209, 209, 209, 0.5);
  border-radius: 3px;
`

const Content = styled.div`
  height: 100%;
`

const Label = styled.h1`
  display: -webkit-box;
  overflow: hidden;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  line-height: 25px;
  height: 50px;
  margin-bottom: 10px;
`

const Footer = styled.div`
  display: flex;
  justify-content: space-around;
  flex-shrink: 0;
`
