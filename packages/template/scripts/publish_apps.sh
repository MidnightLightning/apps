#!/bin/bash
case "$1" in
        staging)
            ENV="staging"
            APM_IPFS_RPC="http://localhost:5001"
            echo environment: $ENV
            ;;
        production)
            ENV="production"
            APM_IPFS_RPC="http://localhost:5001"
            echo environment: $ENV
            ;;
        docker)
            ENV="docker"
            APM_IPFS_RPC="http://ipfs:5001"
            echo environment: $ENV
            ;;
        *)
            ENV="default"
            APM_IPFS_RPC="http://localhost:5001"
            echo environment: $ENV
            ;;
esac

pushd ../distribution
aragon apm publish major --environment $ENV --ipfs-check false --apm.ipfs.rpc $APM_IPFS_RPC
popd

pushd ../hamburger
aragon apm publish major --environment $ENV --ipfs-check false --apm.ipfs.rpc $APM_IPFS_RPC
popd

pushd ../karma-cap-voting
aragon apm publish major --environment $ENV --ipfs-check false --apm.ipfs.rpc $APM_IPFS_RPC
popd

pushd ../registry
aragon apm publish major --environment $ENV --ipfs-check false --apm.ipfs.rpc $APM_IPFS_RPC
popd

pushd ../tipping
aragon apm publish major --environment $ENV --ipfs-check false --apm.ipfs.rpc $APM_IPFS_RPC
popd

pushd ../token-manager
aragon apm publish major --environment $ENV --ipfs-check false --apm.ipfs.rpc $APM_IPFS_RPC
popd
