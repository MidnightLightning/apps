#!/bin/bash
case "$1" in
        staging)
            ENV="staging"
            ENS=0xe7410170f87102DF0055eB195163A03B7F2Bff4A
            RESOLVER=0xB14fdEe4391732eA9d2267054EAd2084684C0aD8
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            ENS=0x5f6f7e8cc7346a11ca2def8f827b7a0b612c56a1
            RESOLVER=0x0000000000000000000000000000000000000000                 # passing 0x will cause default to resolver.eth. basically only need to pass for rinkeby
            echo environment: $ENV
            ;;
esac

cd ~/Projects/daonuts/apps/packages/template
aragon deploy AppInstaller --init $ENS @ARAGON_ENS $RESOLVER --environment $ENV

echo What is the deployed AppInstaller address?
read APP_INSTALLER
export APP_INSTALLER
