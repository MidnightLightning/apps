#!/bin/bash

case "$1" in
        staging)
            ENV="staging"
            NETWORK="rinkeby"
            ENS=0xe7410170f87102DF0055eB195163A03B7F2Bff4A
            ARAGON_ENS=0x98df287b6c145399aaa709692c8d308357bc085d
            APM_ROOT_NODE=0xbf6f73e6e925e595025f4fb0eec5a23cabd74a7a9b0d1f3e5bc88b44fa02e728  # namehash('open.aragonpm.eth')
            RESOLVER=0xB14fdEe4391732eA9d2267054EAd2084684C0aD8
            ROOT_NODE=0xc1d51a8fe29a6c7a959dfdb52aa056a1148543c1250c273d18e4b00585b6f039
            TLD="test"
            TEMPLATE_APM="daonuts-template.open.aragonpm.eth"
            echo environment: $ENV
            ;;
        production)
            ENV="production"
            NETWORK="mainnet"
            ENS=0x314159265dd8dbb310642f98f50c066173c1259b
            ARAGON_ENS=$ENS
            APM_ROOT_NODE=0xbf6f73e6e925e595025f4fb0eec5a23cabd74a7a9b0d1f3e5bc88b44fa02e728  # namehash('open.aragonpm.eth')
            ROOT_NODE=0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e
            TLD="eth"
            TEMPLATE_APM="daonuts-template.open.aragonpm.eth"
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            NETWORK="development"
            ENS=0x5f6f7e8cc7346a11ca2def8f827b7a0b612c56a1
            ARAGON_ENS=$ENS
            APM_ROOT_NODE=0x9065c3e7f7b7ef1ef4e53d2d0b8e0cef02874ab020c1ece79d5f0d3d0111c0ba  # namehash('aragonpm.eth')
            ROOT_NODE=0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e
            TLD="eth"
            TEMPLATE_APM="daonuts-template"
            echo environment: $ENV
            ;;
esac

export REG_ROOT=0x69468d76e397ff6a3e5f0b6bb4912940f2cb1bc46f2a2997cefc3d92b5bdea29
export DIST_ROOT=0x3c29d24944d02f47733fa7e5e198a3e697566f69725817b5e549ab217413a35b
export ROOT_NAME="daonuts"
export ARAGON_ENS
export APM_ROOT_NODE
export TLD
export ENS
export RESOLVER
export ROOT_NODE

cd ~/Projects/daonuts/apps/packages/template

# echo "claim '$ROOT_NAME.$TLD'"
# truffle exec --network $NETWORK scripts/ownRootNode.js

# echo "set resolver"
# truffle exec --network $NETWORK scripts/setResolver.js

. ./scripts/deploy_installer.sh
echo APP_INSTALLER_01: $APP_INSTALLER_01
echo APP_INSTALLER_02: $APP_INSTALLER_02

TEMPLATE=$(aragon deploy Template --init $ARAGON_ENS --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Successfully deployed") print $7 }')
echo Deployed Template to $TEMPLATE
aragon apm publish major $TEMPLATE --environment $ENV

export DAO=$(dao new --template $TEMPLATE_APM --fn newInstance --fn-args $APP_INSTALLER_01 --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Created DAO:") print $5 }')
echo Created DAO: $DAO

echo "install apps 01"
truffle exec --network $NETWORK scripts/installApps01.js

dao apps --all $DAO --environment $ENV

echo What is the currencyManager address?
read CURRENCY_MANAGER
export CURRENCY_MANAGER

echo What is the karmaManager address?
read KARMA_MANAGER
export KARMA_MANAGER

echo What is the voting address?
read VOTING
export VOTING

echo "install apps 02"
truffle exec --network $NETWORK scripts/installApps02.js

dao apps $DAO --environment $ENV

echo What is the registry address?
read REGISTRY
export REGISTRY

echo "transfer root node"
truffle exec --network $NETWORK scripts/transferRootNode.js

echo opening http://localhost:3000/\#/$DAO
xdg-open http://localhost:3000/\#/$DAO
