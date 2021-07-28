package main

import (
	"fmt"
	"log"

	"github.com/indiependente/gocrypto/secretsharing"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	secret := "the secret"
	ss, err := secretsharing.New([]byte(secret), 3, 5)
	if err != nil {
		return err
	}
	ss.ComputeShares()
	shares := ss.Shares()
	for _, s := range shares {
		fmt.Println(s)
	}
	retrievedSecret, err := ss.ComputeSecret(shares)
	if err != nil {
		return err
	}
	fmt.Println(string(retrievedSecret))
	return nil
}
