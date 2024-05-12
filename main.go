package main

import (
	"alchemy/contracts"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	var (
		sk = crypto.ToECDSAUnsafe(common.FromHex(os.Getenv("PKENV")))
	)

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/" + os.Getenv("APIKEY"))
	if err != nil {
		panic(err)
	}

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Unable to get chainid\n")
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(sk, chainid)
	if err != nil {
		log.Fatalf("Unable to get transactopts\n")
	}
	contAddress := common.HexToAddress("0xe7C127f12Dbc68f4114281B82dC1a5DB6515f0A7")
	fmt.Printf("Contract address: %v\n", contAddress)

	contract, err := contracts.NewGetterSetter(contAddress, client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract\n")
	}

	contract.SetUint256(transactOpts, big.NewInt(101))
	if err != nil {
		log.Fatalf("Unable to setUint256\n")
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}

	val, err := contract.GetUint256(callOpts)
	if err != nil {
		log.Fatalf("Unable to get callopts\n")
	}
	fmt.Printf("The value set was: %v\n", val)
}
