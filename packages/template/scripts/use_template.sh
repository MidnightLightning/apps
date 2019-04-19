#!/bin/bash

case "$1" in
        staging)
            ENV="staging"
            NETWORK="rinkeby"
            ENS=0xe7410170f87102DF0055eB195163A03B7F2Bff4A
            ARAGON_ENS=0x98df287b6c145399aaa709692c8d308357bc085d
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
            ROOT_NODE=0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e
            TLD="eth"
            TEMPLATE_APM="daonuts-template"
            echo environment: $ENV
            ;;
esac

export REG_ROOT=0x69468d76e397ff6a3e5f0b6bb4912940f2cb1bc46f2a2997cefc3d92b5bdea29
export DIST_ROOT=0x3c29d24944d02f47733fa7e5e198a3e697566f69725817b5e549ab217413a35b
export ROOT_NAME="daonuts"
export ENS
export ARAGON_ENS
export TLD
export RESOLVER
export ROOT_NODE

cd ~/Projects/daonuts/apps/packages/template

echo "setting resolver (if needed)..."
truffle exec --network $NETWORK scripts/setResolver.js

. ./scripts/deploy_installer.sh
echo AppInstaller: $APP_INSTALLER
TEMPLATE=$(aragon deploy Template --init $ARAGON_ENS --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Successfully deployed") print $7 }')
echo Deployed Template to $TEMPLATE
aragon apm publish major $TEMPLATE --environment $ENV

DAO=$(dao new --template $TEMPLATE_APM --fn newInstance --fn-args $REG_ROOT $DIST_ROOT $APP_INSTALLER --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Created DAO:") print $5 }')
echo Created DAO: $DAO

dao apps $DAO --environment $ENV

echo What is the registry address?
read REGISTRY
export REGISTRY

echo "transfering root node (if needed)..."
truffle exec --network $NETWORK scripts/transferRootNode.js

echo opening http://localhost:3000/\#/$DAO
xdg-open http://localhost:3000/\#/$DAO
