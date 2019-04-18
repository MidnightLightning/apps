#!/bin/bash

cd ~/Projects/daonuts/apps/packages/template
APP_INSTALLER=$(aragon deploy AppInstaller --init $ARAGON_ENS $ENS $RESOLVER $ROOT_NODE --environment $ENV | awk 'NR>1 { if ($3 FS $4 == "Successfully deployed") print $7 }')
echo Deployed AppInstaller to $APP_INSTALLER
export APP_INSTALLER
