package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    registry ".." // for demo
)

func main() {
    client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")

    address := common.HexToAddress("0x711F86c8323C4AFEa09cC6e096B3A34EAE6bEaB4")
    instance, err := registry.NewRegistry(address, client)
    if err != nil {
      log.Fatal(err)
    }

    username := [32]byte{}
    copy(username[:], "dummy01")
    owner, err := instance.UsernameToOwner(nil, username)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println(owner.Hex())

    owner2 := common.HexToAddress("0xb4124cEB3451635DAcedd11767f004d8a28c6eE7")
    username2, err := instance.OwnerToUsername(nil, owner2)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println(string(username2[:]))

}
