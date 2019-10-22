package blockchain

import (
	"container/list"
	"log"
	"time"
)

// Reprezentacja blokchain - lista bloków
type Blockchain struct {
	chain *list.List
}

// Nowy łańcuch
func NewBlockchain(initialBlock *Block) *Blockchain {
	log.Printf("blockChain.go [INFO]: Initilizing New Blockchain")
	chain := list.New()
	chain.PushBack(initialBlock)

	return &Blockchain{chain: chain}
}

// Dodawanie bloku do konkretnej transakcji istniejącego blokchaina
func (bc *Blockchain) AddBlock(t Transaction) error {
	log.Printf("blockChain.go [INFO]: Adding new block to transaktion")
	newBlock := &Block{}
	currentBlock := bc.GetCurrentBlock()

	newBlock.Index = currentBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Transaction = &t
	newBlock.PreviousHash = currentBlock.Hash
	bHash, err := GetBlockHash(*newBlock)
	if err != nil {
		return err
	}
	newBlock.Hash = bHash
	bc.chain.PushBack(newBlock)
	return nil
}

// Zwraca wskaźnik do istniejącego bloku
func (bc *Blockchain) GetCurrentBlock() *Block {
	b := bc.chain.Back().Value.(*Block)
	return b
}
