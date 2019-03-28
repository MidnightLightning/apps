package main

import (
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    hamburger ".." // for demo
)

func main() {
    client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("we have a connection")

    address := common.HexToAddress("0x0858CC5e4B1cdF1f972D44160F341eCfFFe030ab")
    instance, err := hamburger.NewHamburger(address, client)
    if err != nil {
      log.Fatal(err)
    }

    // banner asset id = 0
    asset := big.NewInt(0)
    // what account owns the banner?
    owner, err := instance.OwnerOf(nil, asset)
    if err != nil {
      log.Fatal(err)
    }

    // print owner's account
    fmt.Println(owner.Hex())
}
