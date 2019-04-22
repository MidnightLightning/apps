#!/bin/bash

# cd ~/Projects/daonuts/apps/packages/template
echo using ARAGON_ENS=$ARAGON_ENS
echo using ENS=$ENS
echo using ROOT_NODE=$ROOT_NODE
echo using APM_ROOT_NODE=$APM_ROOT_NODE
APP_INSTALLER_01=$(aragon deploy AppInstaller01 --init $ARAGON_ENS $APM_ROOT_NODE $ENS $ROOT_NODE --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Successfully deployed") print $7 }')
echo Deployed AppInstaller01 to $APP_INSTALLER_01
export APP_INSTALLER_01

APP_INSTALLER_02=$(aragon deploy AppInstaller02 --init $ARAGON_ENS $APM_ROOT_NODE $ENS $ROOT_NODE --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Successfully deployed") print $7 }')
echo Deployed AppInstaller02 to $APP_INSTALLER_02
export APP_INSTALLER_02
