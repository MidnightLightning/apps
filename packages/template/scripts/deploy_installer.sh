#!/bin/bash
case "$1" in
        staging)
            ENV="staging"
            ENS=0xe7410170f87102DF0055eB195163A03B7F2Bff4A
            ARAGON_ENS=0x98df287b6c145399aaa709692c8d308357bc085d
            RESOLVER=0xB14fdEe4391732eA9d2267054EAd2084684C0aD8
            ROOT_NODE=0xc1d51a8fe29a6c7a959dfdb52aa056a1148543c1250c273d18e4b00585b6f039
            echo environment: $ENV
            ;;
        production)
            ENV="production"
            ENS=0x314159265dd8dbb310642f98f50c066173c1259b
            ARAGON_ENS=$ENS
            RESOLVER=0x0000000000000000000000000000000000000000
            ROOT_NODE=0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            ENS=0x5f6f7e8cc7346a11ca2def8f827b7a0b612c56a1
            ARAGON_ENS=$ENS
            RESOLVER=0x0000000000000000000000000000000000000000
            ROOT_NODE=0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e
            echo environment: $ENV
            ;;
esac

cd ~/Projects/daonuts/apps/packages/template
aragon deploy AppInstaller --init $ENS $ARAGON_ENS $RESOLVER $ROOT_NODE --environment $ENV

echo What is the deployed AppInstaller address?
read APP_INSTALLER
export APP_INSTALLER
