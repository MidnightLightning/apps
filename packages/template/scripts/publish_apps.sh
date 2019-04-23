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

pushd ../distribution
aragon apm publish major --environment $ENV
popd

pushd ../hamburger
aragon apm publish major --environment $ENV
popd

pushd ../karma-cap-voting
aragon apm publish major --environment $ENV
popd

pushd ../registry
aragon apm publish major --environment $ENV
popd

pushd ../tipping
aragon apm publish major --environment $ENV
popd

pushd ../token-manager
aragon apm publish major --environment $ENV
popd
