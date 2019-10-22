package main

import (
	"GoExercises/BlockChainImitation/blockchain"
	"bytes"
	"log"
	"math/rand"
)

func main() {
	log.Printf("main.go [INFO]: Blockchain Imitation Application")
	var min = 5
	var max = 50

	// Tworzenie bloku pierwotnego
	bc := blockchain.NewBlockchain(blockchain.NewGenesisBlock())

	// Dodanie kilku transakcji
	for i := 0; i < 10; i++ {
		r1 := rand.Intn(max-min) + min
		r2 := rand.Intn(max-min) + min
		r3 := rand.Intn(max-min) + min
		bc.AddBlock(*blockchain.NewTransaction([]byte{byte(i), byte(r1), byte(r2), byte(r3)}))
	}

	blockChainHash := bc.GetCurrentBlock().Hash

	bc.AddBlock(*blockchain.NewTransaction([]byte{4, 5, 6}))
	prevBlockChainHash := bc.GetCurrentBlock().PreviousHash
	newBlockChainHash := bc.GetCurrentBlock().Hash

	log.Printf("Verification, please wait...")
	var check = bytes.Compare(blockChainHash, prevBlockChainHash)

	if check == 0 {
		log.Printf("Success, hashes are equal and new hash is:")
		log.Println(newBlockChainHash)
	} else {
		log.Printf("Something went wrong, hashes are not equal")
	}
}
