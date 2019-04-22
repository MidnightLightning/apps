#!/bin/bash

# cd ~/Projects/daonuts/apps/packages/template

case "$1" in
        staging)
            ENV="staging"
            NETWORK="rinkeby"
            ENS=0xe7410170f87102DF0055eB195163A03B7F2Bff4A
            ARAGON_ENS=0x98df287b6c145399aaa709692c8d308357bc085d
            RESOLVER=0xB14fdEe4391732eA9d2267054EAd2084684C0aD8
            ROOT_NODE=0xc1d51a8fe29a6c7a959dfdb52aa056a1148543c1250c273d18e4b00585b6f039
            TLD="test"
            APM_ROOT_NAME="open.aragonpm.eth"
            echo environment: $ENV
            ;;
        production)
            ENV="production"
            NETWORK="mainnet"
            ENS=0x314159265dd8dbb310642f98f50c066173c1259b
            ARAGON_ENS=$ENS
            ROOT_NODE=0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e
            TLD="eth"
            APM_ROOT_NAME="open.aragonpm.eth"
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            NETWORK="development"
            ENS=0x5f6f7e8cc7346a11ca2def8f827b7a0b612c56a1
            ARAGON_ENS=$ENS
            ROOT_NODE=0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e
            TLD="eth"
            APM_ROOT_NAME="aragonpm.eth"
            echo environment: $ENV
            ;;
esac

export REG_ROOT=0x69468d76e397ff6a3e5f0b6bb4912940f2cb1bc46f2a2997cefc3d92b5bdea29
export DIST_ROOT=0x3c29d24944d02f47733fa7e5e198a3e697566f69725817b5e549ab217413a35b
export ROOT_NAME="daonuts"
export ENV
export NETWORK
export ARAGON_ENS
export ENS
export ROOT_NODE
export TLD
export APM_ROOT_NAME

# echo "set resolver"
# truffle exec --network $NETWORK scripts/setResolver.js

export DAO=$(dao new --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Created DAO:") print $5 }')
echo Created DAO=$DAO

export APP_INSTALLER=$(aragon deploy AppInstaller --init $ARAGON_ENS $ENS $ROOT_NODE --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Successfully deployed") print $7 }')
if [ -z "$APP_INSTALLER" ]; then
  echo AppInstaller deploy failed
  echo re-running "aragon deploy AppInstaller --init $ARAGON_ENS $ENS $ROOT_NODE --environment $ENV" to show error
  aragon deploy AppInstaller --init $ARAGON_ENS $ENS $ROOT_NODE --environment $ENV
  exit 1
fi
echo Deployed APP_INSTALLER=$APP_INSTALLER

export PERMISSION_SETTER=$(aragon deploy PermissionSetter --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Successfully deployed") print $7 }')
if [ -z "$PERMISSION_SETTER" ]; then
  echo PermissionSetter deploy failed
  echo re-running "aragon deploy PermissionSetter --environment $ENV" to show error
  aragon deploy PermissionSetter --environment $ENV
  exit 1
fi
echo Deployed PERMISSION_SETTER=$PERMISSION_SETTER

dao acl grant $DAO $DAO APP_MANAGER_ROLE $APP_INSTALLER --environment $ENV

export ACL=$(dao apps $DAO --environment $ENV | awk 'NR>1 { if ($2 == "acl") print $4 }')
if [ -z "$ACL" ]; then
  echo ACL proxy address not found
  dao apps $DAO --environment $ENV
  echo What is the ACL Proxy address?
  read ACL
  export ACL
else
  echo Found ACL=$ACL
fi

dao acl grant $DAO $ACL CREATE_PERMISSIONS_ROLE $PERMISSION_SETTER --environment $ENV

echo "install apps"
truffle exec --network $NETWORK scripts/installApps.js

dao acl revoke $DAO $DAO APP_MANAGER_ROLE $APP_INSTALLER --environment $ENV
dao acl revoke $DAO $ACL CREATE_PERMISSIONS_ROLE $PERMISSION_SETTER --environment $ENV

echo opening http://localhost:3000/\#/$DAO
xdg-open http://localhost:3000/\#/$DAO
