As part of reviewing using Aragon to base this project on, I've [deployed to rinkeby an Aragon based dao](https://rinkeby.aragon.org/#/0xeDC1C8aAfb0d4099F8164e3cDE358874C53697bD/) that (ended up) going much of the way to providing a daonuts mvp. It features the following:

* a transferable "commerce" token
* a non-transferable "karma" token
* a customised voting app that employs [KarmaCapVoting](https://www.reddit.com/r/daonuts/comments/ao5umv/why_weight_minrep_token_for_community_governance/) using the commerce and karma tokens. [code is a modified version](https://github.com/daonuts/daonuts-aragon-apps/tree/master/apps/voting) of the [Aragon voting module](https://github.com/aragon/aragon-apps/tree/master/apps/voting).
* a distribution app that accepts a merkle root by dao vote. once accepted users included in the distribution can claim their award. code is custom but based largely on the [recdao merkle validation process](https://github.com/recdao/contracts/blob/master/contracts/RECDAO.sol).
* a [dummy distribution](https://docs.google.com/spreadsheets/d/1Re-eCfBayd67LnJKeTa5GDwdUx6X4r7w5nL3gLQ8lBY/edit#gid=1199924304) that includes users who registered with recdao (~700 ethtraders). if you control an address in that list you could claim a reward in the dao. alternatively, include an address here, or dm me, and I can whip up and include it in new distribution.

**note** - if you visit the dao and only see `Token Manager` and not `Distribution` and `Voting` in the left menu then go to settigns and replace the existing ipfs gateway (`https://ipfs.eth.aragon.network/ipfs`) with `https://gateway.ipfs.io/ipfs/`.


# What is this?

There are different approaches to development that the daonuts project could take, but in my opinion an important one is whether to build on top of an existing "DAO framework" like Aragon or DAOstack. While there are some drawbacks to utilising a framework there is also expected to be some benefit. Generally, a framework may provide:

* base set of audited contracts
* access control system
* modular approach
* ui components

The main drawbacks I can see are:

* the hurdle to understand the framework in the first place
* extraneous code and complexity (larger bug surface, more expensive gas-wise, etc.)

I would say after working with Aragon throughout last week I am largely over the initial hurdle (and therefore could assist others there). In addition, I think the complexity is warranted and would likely need to be recreated with a bespoke development approach - particularly the ACL (permissioning system). I would also say building on top of a set of audited contracts, and modifying or inheriting from ones when customisation is needed, should increase confidence in the eventual system's security - less *new* code is good!

I'll admit that when I looked at Aragon about a year ago I dismissed it due to perceived complexity and because I did not cross of over that initial hurdle quickly enough. This time around I feel completely differently and would conclude that Aragon, and even the work completed on it so far, would be an excellent foundation to build daonuts on.


### Claiming an award

If you control an address in [this list](https://docs.google.com/spreadsheets/d/1Re-eCfBayd67LnJKeTa5GDwdUx6X4r7w5nL3gLQ8lBY/edit#gid=1199924304), you can claim an award in the existing distribution (merle root 0x17c5...). To do so:

1. (switch metamask to rinkeby)
1. copy all the cells in the corresponding row between "root" and "proof" collumns.
1. paste into a text editor and replace the tabs with spaces (sorry for needing this, thought i had that fixed)
1. copy that and paste into the lower, longer, "Claim Award" field
1. line should look something like `0x17c544697d9dafc3d57efd58724d1ab5ac411d466afb8ee429dc460264b083e7 0x95D9bED31423eb7d5B68511E0352Eae39a3CDD20 1500000 ["0x085142a84ec50ddef3e425b437745f8a3ea2195beb3681d8a8c568c76b10ca64","0xbb175278ee5cb7bab6739d98761148e9d0906fb5fb824843f3ef8dca2e6bb579","0x37a038b2390a25183d04c443fe08b6bdbbfd3bc310a25f1dc88e8a395aad4af9","0x904bf3dbd975d6b5474a0375ddde10a80d99b2509669d87a51daaa38f6e08c5f","0xe6af99a84489783659ffc9a3c92718f1103aa1ac2ad05db335ccf9c3998f8080","0xd4518abce405b99b9f2ab9741b8da4d1bf5baf6ab74204674960fcfd742f17e5","0x258e892b26971530b9ba966d70eef3b85fed48e4b356fe4a4a8506d4fda390a1","0xd694e42d5432fc6e532b7ab05adc09d5453cd3ee4f697627491dea9be82e01db","0x9e1bd4ed453b6e6349a5e53e1a66845ba7142244e49276ac5d76f602d52b73b1","0xfbf5c865e3491d49654c4f64717090fa5f7e44d2da8beca05dfe6c4c29cffa7e"]`
1. press "Claim" and the sidebar should open prompting to submit the tx. you will also need to sign with metamask or whatever web3/rinkeby provider you're using.

**note 2!** - your token/karma balance will still display as 0 because I did not account for decimal places. all balances will be off by 10^18 so the ui rounds them to 0 (but they do exist and relative weights mean they still work for voting).

### What's next?

* distribute development work between those interested to contribute
* how should tipping work for the mvp?
* improve the ui/ux of the distribution module
* document the api for use by Reddit devs
