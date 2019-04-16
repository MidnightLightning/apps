#!/bin/bash

REG_ROOT=0x69468d76e397ff6a3e5f0b6bb4912940f2cb1bc46f2a2997cefc3d92b5bdea29
DIST_ROOT=0x3c29d24944d02f47733fa7e5e198a3e697566f69725817b5e549ab217413a35b
ROOT_NAME="daonuts"

case "$1" in
        staging)
            ENV="staging"
            NETWORK="rinkeby"
            ENS=0x98df287b6c145399aaa709692c8d308357bc085d
            OWNER=0x7b6C819e9db25c302A9adD821361bB95524023D7
            TLD="test"
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            NETWORK="development"
            ENS=0x5f6f7e8cc7346a11ca2def8f827b7a0b612c56a1
            OWNER=0xb4124cEB3451635DAcedd11767f004d8a28c6eE7
            TLD="eth"
            echo environment: $ENV
            ;;
esac

export REG_ROOT
export DIST_ROOT
export ENS
export OWNER
export TLD
export ROOT_NAME

cd ~/Projects/daonuts/apps/packages/template
. ./scripts/deploy_installer.sh
echo AppInstaller: $APP_INSTALLER
aragon deploy Template --init @ARAGON_ENS --environment $ENV

echo What is the deployed Template address?
read TEMPLATE

aragon apm publish major $TEMPLATE --environment $ENV

dao new --template daonuts-template --fn newInstance --fn-args $REG_ROOT $DIST_ROOT $APP_INSTALLER --environment $ENV

echo What is the created DAO address?
read DAO

dao apps $DAO --environment $ENV

echo What is the registry address?
read REGISTRY
export REGISTRY

truffle exec --network $NETWORK scripts/transferRootNode.js

echo opening http://localhost:3000/\#/$DAO
xdg-open http://localhost:3000/\#/$DAO
