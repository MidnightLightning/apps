import React from 'react'
import {
  AragonApp,
  Button,
  Text,
  Field,
  TextInput,
  SidePanel,
  theme,

  observe
} from '@aragon/ui'

import styled from 'styled-components'
import color from 'onecolor'

export default class Asset extends React.Component {

  state = {newPrice: 0, creditAmount: 1234, newData: ''}

  openOwnerControls = () => {
    this.props.usePanel({title:"Owner controls", child: OwnerControls, childProps: this.props})
  }

  openBuyControls = () => {
    this.props.usePanel({title:`Buy asset #${this.props.id}`, child: BuyControls, childProps: this.props})
  }

  delete = () => {
    this.props.app.burn(this.props.id)
  }

  render () {
    return (
      <Card>
        <Content>
          <Label>
            <Text color={theme.textTertiary}>#{this.props.id} </Text>
            <span>{web3.toAscii(this.props.name)}</span>
          </Label>
          <Label>
            <Text>{this.props.isOwner ? "You own this asset" : `Owner is: ${this.props.owner.slice(0,10)}...`}</Text>
          </Label>

          <p>tax: {this.props.tax}% / day</p>
          <p>price: {web3.toBigNumber(this.props.price).div("1e+18").toFixed()}</p>
          <p>balance: {web3.toBigNumber(this.props.balance).div("1e+18").toFixed()}</p>
          <p>data: {this.props.data}</p>
          <p>reg required: {this.props.requireReg}</p>
        </Content>
        <Footer>
          <Button mode="outline" onClick={this.delete} emphasis="negative">Delete</Button>
          {this.props.isOwner ?
            <Button mode="strong" onClick={this.openOwnerControls}>Show controls</Button>
            :
            <Button mode="strong" onClick={this.openBuyControls} emphasis="positive">Buy</Button>
          }
        </Footer>
        <Footer>
        </Footer>
      </Card>
    )
  }
}

class OwnerControls extends React.Component {

  state = {newPrice: 0, creditAmount: 1234, newData: ''}

  handleChange = (event) => {
    this.setState({ [event.target.id]: event.target.value });
  }

  presetCredit = (evebt) => {
    let price = web3.toBigNumber(this.props.price).div("1e+18").toNumber()
    switch(event.target.id) {
      case "1-week":
        return this.setState({creditAmount: Math.ceil(1 * 7 * parseInt(this.props.tax) * price / 100)})
      case "2-weeks":
        return this.setState({creditAmount: Math.ceil(2 * 7 * parseInt(this.props.tax) * price / 100)})
      case "4-weeks":
        return this.setState({creditAmount: Math.ceil(4 * 7 * parseInt(this.props.tax) * price / 100)})
    }
  }

  changePrice = () => {
    let newPrice = web3.toBigNumber(this.state.newPrice).mul("1e+18")
    console.log("changePrice", newPrice.toFixed())

    this.props.app.setPrice(this.props.id, newPrice.toFixed())
  }

  changeData = () => {
    this.props.app.setData(this.props.id, this.state.newData)
  }

  credit = async () => {
    let tokenAddress = await this.props.app.call('currency').toPromise()
    let credit = web3.toBigNumber(this.state.creditAmount).mul("1e+18")
    let intentParams = {
      token: { address: tokenAddress, value: credit.toFixed() }
      // gas: 2000000
    }

    let onlyIfSelfOwned = true;

    console.log("credit", credit.toFixed())

    this.props.app.credit(this.props.id, credit.toFixed(), onlyIfSelfOwned, intentParams)
  }

  render(){
    return (
      <React.Fragment>
        <Field label="Change price:">
          <TextInput id="newPrice" type="number" value={this.state.newPrice} onChange={this.handleChange} />
          <Button mode="outline" onClick={this.changePrice}>Change price</Button>
        </Field>
        <Field label="Change data:">
          <TextInput id="newData" placeholder="https://some_image_url" value={this.state.newData} onChange={this.handleChange} />
          <Button mode="outline" onClick={this.changeData}>Change data</Button>
        </Field>
        <Field label="Add to balance:">
          <div>
            <Button mode="outline" size="small" id="1-week" onClick={this.presetCredit}>1 week</Button>
            <Button mode="outline" size="small" id="2-weeks" onClick={this.presetCredit}>2 weeks</Button>
            <Button mode="outline" size="small" id="4-weeks" onClick={this.presetCredit}>4 weeks</Button>
          </div>
          <TextInput id="creditAmount" type="number" value={this.state.creditAmount} onChange={this.handleChange} />
          <Button mode="outline" onClick={this.credit}>Credit</Button>
        </Field>
      </React.Fragment>
    )
  }
}

class BuyControls extends React.Component {

  state = {newPrice: 0, creditAmount: 0, newData: ''}

  handleNewPrice = (event) => {
    let value = parseInt(event.target.value)
    this.setState({
      newPrice: value,
      creditAmount: Math.ceil(1 * 7 * parseInt(this.props.tax) * value / 100)
    })
  }

  handleNewData = (event) => this.setState({newData: event.target.value})

  presetCredit = (evebt) => {
    switch(event.target.id) {
      case "2-weeks":
        return this.setState({creditAmount: Math.ceil(2 * 7 * parseInt(this.props.tax) * this.state.newPrice / 100)})
      case "4-weeks":
        return this.setState({creditAmount: Math.ceil(4 * 7 * parseInt(this.props.tax) * this.state.newPrice / 100)})
    }
  }

  buy = async () => {
    let tokenAddress = await this.props.app.call('currency').toPromise()

    let newPrice = web3.toBigNumber(this.state.newPrice).mul("1e+18")
    let newData = this.state.newData
    let credit = web3.toBigNumber(this.state.creditAmount).mul("1e+18")
    let value = web3.toBigNumber(this.props.price).add(credit)
    console.log("buy", value, value.toFixed(), credit.toFixed(), newPrice.toFixed())

    let intentParams = {
      token: { address: tokenAddress, value: value.toFixed() }
      // gas: 2000000
    }

    this.props.app.buy(this.props.id, newPrice.toFixed(), newData, credit.toFixed(), intentParams)
  }

  render(){
    return (
      <React.Fragment>
        <Field label="New price:">
          <TextInput id="newPrice" type="number" value={this.state.newPrice} onChange={this.handleNewPrice} />
        </Field>
        <Field label="New data: (can change later)">
          <TextInput id="newData" placeholder="https://some_image_url" value={this.state.newData} onChange={this.handleNewData} />
        </Field>
        <Field label="Start balance: (min=1 day tax, can add to later)">
          <div>
            <Button mode="outline" size="small" id="2-weeks" onClick={this.presetCredit}>2 weeks</Button>
            <Button mode="outline" size="small" id="4-weeks" onClick={this.presetCredit}>4 weeks</Button>
          </div>
          <TextInput id="creditAmount" type="number" value={this.state.creditAmount} onChange={this.handleChange} />
        </Field>
        <Button mode="outline" onClick={this.buy}>Buy</Button>
      </React.Fragment>
    )
  }
}

const SecondaryButton = styled(Button).attrs({
  mode: 'secondary',
  compact: true,
})`
  background: ${color(theme.secondaryBackground)
    .alpha(0.8)
    .cssa()};
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
