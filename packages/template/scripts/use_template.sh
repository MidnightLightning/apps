#!/bin/bash

REG_ROOT=0xbbdacbe4195b9287f120e672eebc85fa1877ba9a0db7b1a3414d5846f08dc160
DIST_ROOT=0x4e42be0076fb77392a609d01b2b30f533e18a9cc1adafb1c534d40bb7bc6472c

case "$1" in
        staging)
            ENV="staging"
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            echo environment: $ENV
            ;;
esac

cd ~/Projects/daonuts/apps/packages/template
aragon deploy AppInstaller --environment $ENV

echo What is the deployed AppInstaller address?
read APP_INSTALLER

aragon deploy Template --init @ARAGON_ENS --environment $ENV

echo What is the deployed Template address?
read TEMPLATE

aragon apm publish major $TEMPLATE --environment $ENV

dao new --template daonuts-template --fn newInstance --fn-args $REG_ROOT $DIST_ROOT $APP_INSTALLER --environment $ENV

echo What is the created DAO address?
read DAO

echo opening http://localhost:3000/\#/$DAO
xdg-open http://localhost:3000/\#/$DAO
