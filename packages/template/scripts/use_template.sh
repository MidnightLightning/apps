#!/bin/bash

REG_ROOT=0xbbdacbe4195b9287f120e672eebc85fa1877ba9a0db7b1a3414d5846f08dc160
DIST_ROOT=0x4e42be0076fb77392a609d01b2b30f533e18a9cc1adafb1c534d40bb7bc6472c
ENS=0x5f6f7e8cc7346a11ca2def8f827b7a0b612c56a1
OWNER=0xb4124cEB3451635DAcedd11767f004d8a28c6eE7

export REG_ROOT
export DIST_ROOT
export ENS
export OWNER

case "$1" in
        staging)
            ENV="staging"
            NETWORK="rinkeby"
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            NETWORK="development"
            echo environment: $ENV
            ;;
esac

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
