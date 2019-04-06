#!/bin/bash
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
export APP_INSTALLER
