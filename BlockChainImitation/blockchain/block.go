package blockchain

import (
	"bytes"
	"crypto/sha256"
	"log"
	"time"
)

// Block reprezentuje pojedyńczą transakcje
type Block struct {
	Index        uint64
	Timestamp    string
	Transaction  *Transaction
	Hash         []byte
	PreviousHash []byte
}

// Tworzenie nowego bloku
func NewGenesisBlock() *Block {
	log.Printf("block.go [INFO]: Creating new block")
	initialBlock := Block{
		Index:       uint64(0),
		Timestamp:   time.Now().String(),
		Transaction: NewTransaction([]byte{}),
	}

	bHash, _ := GetBlockHash(initialBlock)
	initialBlock.Hash = bHash
	initialBlock.PreviousHash = nil

	return &initialBlock
}

// Pobieranie hasha bloku
func GetBlockHash(b Block) ([]byte, error) {
	buf := bytes.NewBufferString("")
	buf.WriteString(string(b.Index))
	buf.WriteString(b.Timestamp)
	buf.Write(b.Transaction.Hash)
	buf.Write(b.PreviousHash)

	h := sha256.New()
	h.Write(buf.Bytes())

	return h.Sum(nil), nil
}
