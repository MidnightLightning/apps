#!/bin/bash

REG_ROOT=0xbbdacbe4195b9287f120e672eebc85fa1877ba9a0db7b1a3414d5846f08dc160
DIST_ROOT=0x4e42be0076fb77392a609d01b2b30f533e18a9cc1adafb1c534d40bb7bc6472c

# export ME=0xb4124cEB3451635DAcedd11767f004d8a28c6eE7
# export ME=0xF606d6a3De1592F834FD905781c85b01F5e5f995
# export ME=0xA38f8165a4e512FA4D2f9930C0B1b5DaD0f1820D
export ME=0x7b6C819e9db25c302A9adD821361bB95524023D7
# export ME=0xB1e442E09512b1792B88498ef4C9537F68621C5f
echo Deployer address is: $ME

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

case "$2" in
        1)
            echo '(1) Creating base DAO...'
            dao new --environment $ENV
            echo What is the deployed DAO address?
            read DAO
            export DAO
            ;&

        2)
            echo '(2) Creating commerce token...'
            dao token new Commerce DAOC 18 true --environment $ENV
            echo What is the deployed Commerce token address?
            read DAOC
            export DAOC
            ;&

        3)
            echo '(3) Installing commerce token manager...'
            dao install $DAO token-manager --app-init none --environment $ENV
            dao apps $DAO --all --environment $ENV
            echo What is the the deployed token manager address?
            read DAOCM
            export DAOCM
            ;&

        4)
            echo '(4) Set token controller...'
            dao token change-controller $DAOC $DAOCM --environment $ENV
            ;&

        5)
            echo '(5) Assign token manager MINT_ROLE to $ME...'
            dao acl create $DAO $DAOCM MINT_ROLE $ME $ME --environment $ENV
            ;&

        6)
            echo '(6) Inititializing commerce token managers...'
            dao exec $DAO $DAOCM initialize $DAOC true 0 --environment $ENV
            ;&

        7)
            echo '(7) Assigning 1 DAOC to $ME...'
            dao exec $DAO $DAOCM mint $ME 1 --environment $ENV
            ;&

        8)
            echo '(8) Creating karma token...'
            dao token new Karma DAOK 18 false --environment $ENV
            echo What is the deployed Karma token address?
            read DAOK
            export DAOK
            ;&

        9)
            echo '(9) Installing karma token manager...'
            dao install $DAO token-manager --app-init none --environment $ENV
            dao apps $DAO --all --environment $ENV
            echo What is the the karma token manager address?
            read DAOKM
            export DAOKM
            ;&

        10)
            echo '(10) Set karma token controller...'
            dao token change-controller $DAOK $DAOKM --environment $ENV
            ;&

        11)
            echo '(11) Assign karma token manager MINT_ROLE to $ME...'
            dao acl create $DAO $DAOKM MINT_ROLE $ME $ME --environment $ENV
            ;&

        12)
            echo '(12) Inititializing karma token managers...'
            dao exec $DAO $DAOKM initialize $DAOK false 0 --environment $ENV
            ;&

        13)
            echo '(13) Assigning 1 DAOK to $ME...'
            dao exec $DAO $DAOKM mint $ME 1 --environment $ENV
            ;&

        14)
            echo '(14) Installing KarmaCapVoting 50%, 10% quorum, 1 day...'
            dao install $DAO daonuts-karma-cap-voting.open.aragonpm.eth --app-init-args $DAOC $DAOK 500000000000000000 100000000000000000 86400 --environment $ENV
            # find app proxy address
            dao apps $DAO --all --environment $ENV
            echo What is the the deployed voting address?
            read VOTING
            export VOTING
            ;&

        15)
            echo '(15) Create permission...'
            dao acl create $DAO $VOTING CREATE_VOTES_ROLE $DAOKM $VOTING --environment $ENV
            ;&

        16)
            echo '(16) Installing registry...'
            dao install $DAO daonuts-registry.open.aragonpm.eth --app-init-args $REG_ROOT --environment $ENV
            # find app proxy address
            dao apps $DAO --all --environment $ENV
            echo What is the the deployed registry address?
            read REGISTRY
            export REGISTRY
            ;&

        17)
            echo '(17) Create permission...'
            dao acl create $DAO $REGISTRY START_REGISTRATION_PERIOD $VOTING $VOTING --environment $ENV
            ;&

        18)
            echo '(18) Installing distribution...'
            dao install $DAO daonuts-distribution.open.aragonpm.eth --app-init-args $DAOCM $DAOKM $REGISTRY $DIST_ROOT --environment $ENV
            # find app proxy address
            dao apps $DAO --all --environment $ENV
            echo What is the the deployed distribution address?
            read DISTRIBUTION
            export DISTRIBUTION
            ;&

        19)
            echo '(19) Create permission...'
            dao acl create $DAO $DISTRIBUTION START_DISTRIBUTION $VOTING $VOTING --environment $ENV
            ;&

        20)
            echo '(20) Installing tipping...'
            dao install $DAO daonuts-tipping.open.aragonpm.eth --environment $ENV --app-init-args $DAOC $REGISTRY
            # find app proxy address
            dao apps $DAO --all --environment $ENV
            echo What is the the deployed tipping address?
            read TIPPING
            export TIPPING
            ;&

        21)
            echo '(21) Create permission...'
            dao acl create $DAO $TIPPING NONE $TIPPING $TIPPING --environment $ENV
            ;&

        22)
            echo '(20) Installing hamburger...'
            dao install $DAO daonuts-hamburger.open.aragonpm.eth --environment $ENV --app-init-args $DAOC $REGISTRY
            # find app proxy address
            dao apps $DAO --all --environment $ENV
            echo What is the the deployed hamburger address?
            read HAMBURGER
            export HAMBURGER
            ;&

        23)
            echo '(21) Create permission...'
            dao acl create $DAO $HAMBURGER COMMONS_ROLE $VOTING $VOTING --environment $ENV
            ;&

        24)
            echo '(21) Create permission...'
            dao acl create $DAO $HAMBURGER PURCHASE_ASSET_ROLE $REGISTRY $VOTING --environment $ENV
            ;&

        25)
            echo '(22) Setting up more permissions...'
            dao acl grant $DAO $DAOCM MINT_ROLE $VOTING --environment $ENV
            ;&

        26)
            echo '(23) Setting up more permissions...'
            dao acl grant $DAO $DAOKM MINT_ROLE $VOTING --environment $ENV
            ;&

        27)
            echo '(24) Setting up more permissions...'
            dao acl grant $DAO $DAOCM MINT_ROLE $DISTRIBUTION --environment $ENV
            ;&

        28)
            echo '(25) Setting up more permissions...'
            dao acl grant $DAO $DAOKM MINT_ROLE $DISTRIBUTION --environment $ENV
            ;;


        *)
            echo $"Usage: $0 {staging} $0 {1-25}"
            ;;
esac
