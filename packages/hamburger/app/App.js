import React from 'react'
import {
  AragonApp,
  AppBar,
  AppView,
  BaseStyles,
  Main,
  Button,
  Text,
  Field,
  TextInput,
  SidePanel,
  Checkbox,
  breakpoint,
  observe
} from '@aragon/ui'
import Aragon, { providers } from '@aragon/client'
import styled from 'styled-components'
import Asset from './Asset'

const AppContainer = styled(AragonApp)`
  display: flex;
  align-items: center;
  justify-content: center;
`

export default class App extends React.Component {

  state = {assets: [], username:'', panelOpen: false, panel: {title: "Peach"}}

  lastObservable = {}

  initDone = false

  componentDidUpdate = (prevProps) => {
    if(!this.initDone) {
      this.init()
    }
    if (this.props.observable && !prevProps.observable) {
      this.props.observable.subscribe(o=>{
        if(!o) return
        if(o.assetsCount !== this.lastObservable.assetsCount)
          this.getAssets(o.assetsCount)
        if(this.state.assets.length && o.transfer && o.transfer !== this.lastObservable.transfer)
          this.getAsset(o.transfer.id)
        this.lastObservable = o
      })
    }
  }

  init = async () => {
    let assetsCount = await this.props.app.call('assetsCount').toPromise()
    this.getAssets(assetsCount)
    this.initDone = true
  }

  getAsset = async (id) => {
    console.log("need to update id:", id)
    let assets = this.state.assets.slice(0);
    let idx = assets.findIndex(a=>a.id===id)
    if(idx === -1)
      return
    let asset = await this.props.app.call('assets', id).toPromise()

    assets[idx] = Object.assign(assets[idx], asset);
    this.setState({assets})
  }

  getAssets = async (length) => {

    let assets = await Promise.all(
        Array
          .from({length})
          .map((v,i)=>this.props.app.call('assets', i).toPromise())
      )

    assets.forEach((v,i)=>v.id=i)

    this.setState({assets})
  }

  usePanel = (panel) => {
    this.setState({panelOpen: true, panel})
  }

  closePanel = () => {
    this.setState({panelOpen: false})
  }

  handleMenuPanelOpen = () => {
    this.props.sendMessageToWrapper('menuPanel', true)
  }

  openNewAssetControls = () => {
    this.usePanel({title:"New asset", child: NewAsset, childProps: this.props})
  }

  render () {
    const Child = this.state.panel.child
    return (
      <Main>
        <AppView appBar={<AppBar title="Hamburgers" endContent={<Button mode="strong" onClick={this.openNewAssetControls}>New Asset</Button>} />} >
        <AssetList assets={this.state.assets} app={this.props.app} userAccount={this.props.userAccount} usePanel={this.usePanel} />
        <SidePanel title={this.state.panel.title} opened={this.state.panelOpen} onClose={this.closePanel}>
          {Child && <Child {...this.state.panel.childProps} />}
        </SidePanel>
      </AppView>
    </Main>
    )
  }
}

class NewAsset extends React.Component {

  state = {newName: '', newTax: 1, requireReg: true}

  handleChange = (event) => {
    this.setState({ [event.target.id]: event.target.value });
  }

  add = () => {
    let name = web3.fromAscii(this.state.newName)
    let tax = this.state.newTax
    let requireReg = true
    this.props.app.mint(name, tax, requireReg)
  }

  render(){
    return (
      <React.Fragment>
        <Field label="New asset name:">
          <TextInput id="newName" placeholder="name" value={this.state.newName} onChange={this.handleChange} />
        </Field>
        <Field label="New asset tax (%/day):">
          <TextInput id="newTax" type="number" value={this.state.newTax} onChange={this.handleChange} />
        </Field>
        <Button onClick={this.add}>Add Asset</Button>
      </React.Fragment>
    )
  }
}
// <Field label="Require owner to be registered:">
//   <label>
//     <Checkbox checked={this.state.requireReg} onChange={checked => this.setState({ requireReg: checked })} />
//     Require registration
//   </label>
// </Field>

function AssetList({assets, app, usePanel, userAccount}) {
  const listItems = assets.filter(a=>a.active).map((asset) => <Asset {...asset} app={app} usePanel={usePanel} isOwner={userAccount===asset.owner} />);
  return (
    <Grid>{listItems}</Grid>
  );
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
