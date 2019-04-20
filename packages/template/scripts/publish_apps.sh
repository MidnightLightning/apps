#!/bin/bash
case "$1" in
        staging)
            ENV="staging"
            echo environment: $ENV
            ;;
        production)
            ENV="production"
            echo environment: $ENV
            ;;
        docker)
            ENV="docker"
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            echo environment: $ENV
            ;;
esac

pushd ./packages/distribution
aragon apm publish major --environment $ENV
popd

pushd ./packages/hamburger
aragon apm publish major --environment $ENV
popd

pushd ./packages/karma-cap-voting
aragon apm publish major --environment $ENV
popd

pushd ./packages/registry
aragon apm publish major --environment $ENV
popd

pushd ./packages/tipping
aragon apm publish major --environment $ENV
popd

pushd ./packages/token-manager
aragon apm publish major --environment $ENV
popd
