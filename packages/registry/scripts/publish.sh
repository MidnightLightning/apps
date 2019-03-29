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

case "$2" in
        patch)
            BUMP="patch"
            echo bump: $BUMP
            ;;
        minor)
            BUMP="minor"
            echo bump: $BUMP
            ;;
        *)
            BUMP="major"
            echo bump: $BUMP
            ;;
esac

npm run deploy:local
echo What is the deployed Registry address?
read REGISTRY
aragon apm publish $BUMP $REGISTRY --environment $ENV
