#!/bin/bash
case "$1" in
        staging)
            ENV="staging"
            ENS=0xe7410170f87102DF0055eB195163A03B7F2Bff4A
            RESOLVER=0xB14fdEe4391732eA9d2267054EAd2084684C0aD8
            ROOT_NODE=0xc1d51a8fe29a6c7a959dfdb52aa056a1148543c1250c273d18e4b00585b6f039
            echo environment: $ENV
            ;;
        production)
            ENV="production"
            ENS="@ARAGON_ENS"
            RESOLVER=0x0000000000000000000000000000000000000000
            ROOT_NODE=0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            ENS="@ARAGON_ENS"
            RESOLVER=0x0000000000000000000000000000000000000000
            ROOT_NODE=0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e
            echo environment: $ENV
            ;;
esac

cd ~/Projects/daonuts/apps/packages/template
APP_INSTALLER=$(aragon deploy AppInstaller --init @ARAGON_ENS $ENS $RESOLVER $ROOT_NODE --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Successfully deployed") print $7 }')
echo Deployed AppInstaller to $APP_INSTALLER
export APP_INSTALLER
