export ME=0xb4124cEB3451635DAcedd11767f004d8a28c6eE7

dao new --environment aragon:rinkeby
export DAO=0xeDC1C8aAfb0d4099F8164e3cDE358874C53697bD

dao token new DaonutsCommerce NUTC 18 true --environment aragon:rinkeby
export NUTC=0xe02eda7b7178B5CAAfFd452C63F2B2955770F13f

dao token new DaonutsKarma NUTK 18 false --environment aragon:rinkeby
export NUTK=0x9A0481065453DF7cEaa4D2d698BAc7fe839407Db

// install 2 token managers
dao install $DAO token-manager --app-init none --environment aragon:rinkeby
dao install $DAO token-manager --app-init none --environment aragon:rinkeby

dao apps $DAO --all --environment aragon:rinkeby

export NUTCM=0xa9827F65CE5CD26288c8d3D1b1CfC97be07d196D
export NUTKM=0x89Ed6088002A3Bd974DcC32DDCDa2d11167727B1

dao token change-controller $NUTC $NUTCM --environment aragon:rinkeby
dao token change-controller $NUTK $NUTKM --environment aragon:rinkeby

dao acl create $DAO $NUTCM MINT_ROLE $ME $ME --environment aragon:rinkeby
dao acl create $DAO $NUTKM MINT_ROLE $ME $ME --environment aragon:rinkeby

dao exec $DAO $NUTCM initialize $NUTC true 0 --environment aragon:rinkeby
dao exec $DAO $NUTKM initialize $NUTK false 0 --environment aragon:rinkeby

dao exec $DAO $NUTCM mint $ME 1 --environment aragon:rinkeby
dao exec $DAO $NUTKM mint $ME 1 --environment aragon:rinkeby

// in Voting app directory
aragon apm publish patch --environment staging

// install KarmaCapVoting 50%, 10% quorum, 1 day
dao install $DAO karma-cap-voting.open.aragonpm.eth --app-init-args $NUTC $NUTK 500000000000000000 100000000000000000 86400 --environment aragon:rinkeby
// find app proxy address
dao apps $DAO --all --environment aragon:rinkeby
export VOTING=0x20494640518Fdd4221a05EE67eDbE7198b4eC4c1

// in Distribution app directory
aragon apm publish patch --environment staging
// find app proxy address
dao apps $DAO --all --environment aragon:rinkeby
export DISTRIBUTION=0x9CA1C996d107e0e1B7B682700E60192605c49dc8

// install Distribution
dao install $DAO distribution.open.aragonpm.eth --app-init-args $NUTCM $NUTKM --environment aragon:rinkeby

// permissions
// dao acl create $DAO [app-proxy-addr] [role] [entity] [manager] --environment aragon:rinkeby
// dao acl grant $DAO [app-proxy-addr] [role] [entity] --environment aragon:rinkeby
<!-- Entity	App	Role	Manager -->
dao acl create $DAO $DISTRIBUTION ADD_ROOT $VOTING $VOTING --environment aragon:rinkeby
dao acl create $DAO $VOTING CREATE_VOTES_ROLE $NUTKM $VOTING --environment aragon:rinkeby
dao acl grant $DAO $NUTCM MINT_ROLE $VOTING --environment aragon:rinkeby
dao acl grant $DAO $NUTKM MINT_ROLE $VOTING --environment aragon:rinkeby
dao acl grant $DAO $NUTCM MINT_ROLE $DISTRIBUTION --environment aragon:rinkeby
dao acl grant $DAO $NUTKM MINT_ROLE $DISTRIBUTION --environment aragon:rinkeby


carl@carl-laptop:~/Projects/aragon-apps/apps/voting$ aragon apm publish patch --environment staging
 ✔ Check IPFS
 ✔ Applying version bump (patch)
 ✔ Determine contract address for version
 ✔ Building frontend
 ✔ Prepare files for publishing
 ✔ Generate application artifact
 ✔ Publish karma-cap-voting.open.aragonpm.eth
 ✔ Fetch published repo

 ✔ Successfully published karma-cap-voting.open.aragonpm.eth v0.0.2:
 ℹ Contract address: 0xefCdfAfCe6a6685FFe776f33Cb5D688395D6D4De
 ℹ Content (ipfs): QmUVZqEtp7xEm2NdhN3Lk7hXsAZes28p1S2pUk7LeaKbmx
 ℹ Transaction hash: 0xc543f4ea92a7836111aa64af30f4c3a378fc99ff57d55f607b023f6d4995c003

carl@carl-laptop:~/Projects/distribution$ aragon apm publish patch --environment staging
 ✔ Check IPFS
 ✔ Applying version bump (patch)
 ✔ Determine contract address for version
 ✔ Building frontend
 ✔ Prepare files for publishing
 ✔ Generate application artifact
 ✔ Publish distribution.open.aragonpm.eth
 ✔ Fetch published repo

 ✔ Successfully published distribution.open.aragonpm.eth v0.0.2:
 ℹ Contract address: 0x79172036D804733d26781a7d67Ab543b711E03C7
 ℹ Content (ipfs): QmfLLgKXUyHuPuAwrxhYJs9MZJgAoCrLoKs7xGyzk3PkJk
 ℹ Transaction hash: 0x55d0d2ed74fca673c5ba962c36a8c04dbba11f8e5838884ae371557dc2fe1795
