#!/bin/bash

HASHES=( QmT4pZSLvkmrS9ZbM9hDRirbJHUNJytgmN934i4guJgigK QmaNozbDgMtqfwDJyZamfF4B1bkmNCnWzTefqC4Pmp7KMw QmTj956kbZkpZiP2HqLqD3GFymuqbJP9CdrkAvmSuHZ3xM QmPr1bFSf3Pn3io3GxaEaGyuqBwepFZ6Xi381S6FqgmwuP Qmd4VHsJtMLgE7imXQMnyd52XhJxzve2PmCT3QbyG69DY4 )

for HASH in "${HASHES[@]}"
do
	curl https://ipfs.lol/ipfs/$HASH
	curl https://ipfs.lol/ipfs/$HASH/arapp.json
	curl https://ipfs.lol/ipfs/$HASH/artifact.json
	curl https://ipfs.lol/ipfs/$HASH/manifest.json
	curl https://ipfs.lol/ipfs/$HASH/dist/index.html
done

for HASH in "${HASHES[@]}"
do
	curl https://ipfs.eth.aragon.network/ipfs/$HASH
	curl https://ipfs.eth.aragon.network/ipfs/$HASH/arapp.json
	curl https://ipfs.eth.aragon.network/ipfs/$HASH/artifact.json
	curl https://ipfs.eth.aragon.network/ipfs/$HASH/manifest.json
	curl https://ipfs.eth.aragon.network/ipfs/$HASH/dist/index.html
done
