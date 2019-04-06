const MerkleTree = require("merkle-tree-solidity").default
const utils = require("ethereumjs-util")
const argv = require('yargs').argv
const csv = require('csvtojson')
var json2csv = require('json2csv').parse;
var fs = require("fs");
const BigNumber = require('bignumber.js');
const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));

parseCSV(argv.date);

async function parseCSV(date){
  let recipients = await csv().fromFile(`${__dirname}/../registrations/pre/${date}.csv`)
  // console.log(recipients)
  // recipients.map(r=>{r.award = parseInt(r.award); return r;})    // Important!
  // recipients.map(r=>{r.awardHex = web3.utils.toHex(new BigNumber(r.award)); return r;})    // Important!
  // console.log(recipients[0].award)
  buildTree(recipients, date)
}

function buildTree(recipients, date){
  const recipientHashBuffers = recipients.map(u=>{
    // console.log(typeof u.award)
    let addressBuffer = utils.toBuffer(u.address)
    // let usernameBuffer = utils.setLengthRight(utils.toBuffer(u.username), 32)
    let usernameBuffer = utils.toBuffer(u.username)
    let hashBuffer = utils.keccak256(Buffer.concat([addressBuffer, usernameBuffer]))
    let hash = utils.bufferToHex(hashBuffer)
    // console.log(hash)

    return hashBuffer
  })

  const merkleTree = new MerkleTree(recipientHashBuffers)

  const root = utils.bufferToHex(merkleTree.getRoot())

  recipients = recipients.map((recipient,idx)=>{
    recipient.root = root
    recipient.proof = merkleTree.getProof(recipientHashBuffers[idx]).map(p=>utils.bufferToHex(p))
    return recipient
  })

  // console.log(recipientHashBuffers.map(utils.bufferToHex))
  console.log(`root:`, root)
  // console.log(`recipients:\n`, recipients)
  // for (var i=0;i<recipients.length;i++){
  //   console.log(`${recipients[i].address} hash: ${utils.bufferToHex(recipientHashBuffers[i])}`)
  // }
  console.log(`${recipients[0].address} hash: ${utils.bufferToHex(recipientHashBuffers[0])}`)

  outputCSV(recipients, root)
}

function outputCSV(data, root) {
    var fields = ['address', 'root', 'username', 'proof'];

    try {
      fs.writeFileSync(`${__dirname}/../registrations/post/${root.slice(0,10)}.json`, JSON.stringify(data))
      var result = json2csv(data, { fields, delimiter: `\t` });
      fs.writeFileSync(`${__dirname}/../registrations/post/${root.slice(0,10)}.csv`, result)
    } catch (err) {
      console.error(err);
    }
}
