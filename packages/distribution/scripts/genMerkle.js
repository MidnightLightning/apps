const MerkleTree = require("merkle-tree-solidity").default
const utils = require("ethereumjs-util")
const argv = require('yargs').argv
const setLengthLeft = utils.setLengthLeft
const setLengthRight = utils.setLengthRight
const csv = require('csvtojson')
var json2csv = require('json2csv').parse;
var fs = require("fs");
const BigNumber = require('bignumber.js');
const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
const decimals = new BigNumber("1e+18")

parseCSV(argv.date);

async function parseCSV(date){
  let recipients = await csv().fromFile(`${__dirname}/../distributions/pre/${date}.csv`)
  recipients = recipients.reduce((prev, curr)=>{
    let username = curr.username.replace('u/','')
    let award = new BigNumber(curr.points)
    let existing = prev.find(u=>u.username===username)
    if(existing) existing.award = existing.award.plus(award)
    else prev.push({username,award})
    return prev
  }, [])
  recipients.forEach(r=>{
    r.award = r.award.times(decimals)
    r.awardHex = web3.utils.toHex(r.award)
    r.award = r.award.toFixed()
  })
  buildTree(recipients, date)
}

function buildTree(recipients, date){
  const recipientHashBuffers = recipients.map(u=>{
    // console.log(typeof u.award)
    let usernameBuffer = utils.setLengthRight(utils.toBuffer(u.username), 32)
    let uintBuffer = setLengthLeft(utils.toBuffer(u.awardHex), 32)
    let hashBuffer = utils.keccak256(Buffer.concat([usernameBuffer, uintBuffer]))
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
  console.log(`${recipients[0].username} hash: ${utils.bufferToHex(recipientHashBuffers[0])}`)

  outputCSV(recipients, root)
}

function outputCSV(data, root) {
    var fields = ['root', 'username', 'award', 'proof'];

    try {
      fs.writeFileSync(`${__dirname}/../distributions/post/${root.slice(0,10)}.json`, JSON.stringify(data))
      var result = json2csv(data, { fields, delimiter: `\t` });
      fs.writeFileSync(`${__dirname}/../distributions/post/${root.slice(0,10)}.csv`, result)
    } catch (err) {
      console.error(err);
    }
}
