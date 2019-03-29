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

cd ~/Projects/daonuts/apps/packages/distribution
aragon apm publish major --environment $ENV

cd ~/Projects/daonuts/apps/packages/hamburger
aragon apm publish major --environment $ENV

cd ~/Projects/daonuts/apps/packages/karma-cap-voting
aragon apm publish major --environment $ENV

cd ~/Projects/daonuts/apps/packages/registry
npm run publish:major

cd ~/Projects/daonuts/apps/packages/tipping
aragon apm publish major --environment $ENV

cd ~/Projects/daonuts/apps/packages/token-manager
aragon apm publish major --environment $ENV
