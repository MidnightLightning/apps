#!/bin/bash

cd ~/Projects/daonuts/apps/packages/template
echo ARAGON_ENS=$ARAGON_ENS ENS=$ENS ROOT_NODE=$ROOT_NODE
APP_INSTALLER=$(aragon deploy AppInstaller --init $ARAGON_ENS $ENS $ROOT_NODE --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Successfully deployed") print $7 }')
echo Deployed AppInstaller to $APP_INSTALLER
export APP_INSTALLER
